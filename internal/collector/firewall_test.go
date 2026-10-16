package collector

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AthennaMind/opnsense-exporter/internal/options"
	"github.com/AthennaMind/opnsense-exporter/opnsense"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/promslog"
)

var firewallTestResponses = map[string]string{
	"/api/diagnostics/firewall/pf_statistics/interfaces": `{"interfaces": {"igb0": {
		"references": 7, "cleared": "2026-10-01T10:00:00",
		"in4_pass_packets": 10, "in4_block_packets": 1, "out4_pass_packets": 9, "out4_block_packets": 0,
		"in6_pass_packets": 5, "in6_block_packets": 0, "out6_pass_packets": 4, "out6_block_packets": 0,
		"in4_pass_bytes": 123456789012, "in4_block_bytes": 60, "out4_pass_bytes": 900, "out4_block_bytes": 0,
		"in6_pass_bytes": 500, "in6_block_bytes": 0, "out6_pass_bytes": 400, "out6_block_bytes": 0}}}`,
	"/api/diagnostics/firewall/pf_statistics/info": `{"info": {
		"state-table": {"current-entries": {"total": 512}, "inserts": {"total": 9000, "rate": 1.0}},
		"source-tracking-table": {"current-entries": {"total": 0}},
		"counters": {"match": {"total": 42, "rate": 0.1}},
		"limit-counters": {"max-src-states": {"total": 0, "rate": 0.0}}}}`,
	"/api/diagnostics/firewall/pf_statistics/memory":   `{"memory": {"states": 400000, "src-nodes": 400000}}`,
	"/api/diagnostics/firewall/pf_statistics/timeouts": `{"timeouts": {"tcp.first": "120s", "adaptive.start": "240000 states"}}`,
}

func newFirewallTestCollector(t *testing.T, failPath string) (*firewallCollector, *opnsense.Client) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, ok := firewallTestResponses[r.URL.Path]
		if !ok || r.URL.Path == failPath {
			http.Error(w, "boom", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	client, err := opnsense.NewClient(options.OPNSenseConfig{
		Protocol:  "http",
		Host:      strings.TrimPrefix(server.URL, "http://"),
		APIKey:    "test",
		APISecret: "test",
	}, "test", promslog.NewNopLogger())
	if err != nil {
		t.Fatalf("failed to build client: %v", err)
	}

	c := &firewallCollector{subsystem: FirewallSubsystem}
	c.Register(namespace, "test", promslog.NewNopLogger())
	return c, &client
}

// collectFirewall runs Update and returns the metric values keyed by
// metric name plus the sorted non-instance label values.
func collectFirewall(t *testing.T, c *firewallCollector, client *opnsense.Client) (map[string]float64, *opnsense.APICallError) {
	t.Helper()

	ch := make(chan prometheus.Metric, 1000)
	apiErr := c.Update(client, ch)
	close(ch)

	values := map[string]float64{}
	for m := range ch {
		var pb dto.Metric
		if err := m.Write(&pb); err != nil {
			t.Fatalf("failed to write metric: %v", err)
		}
		desc := m.Desc().String()
		name := desc[strings.Index(desc, `fqName: "`)+len(`fqName: "`):]
		name = name[:strings.Index(name, `"`)]

		key := name
		for _, l := range pb.GetLabel() {
			if l.GetName() != instanceLabelName {
				key += "|" + l.GetValue()
			}
		}

		switch {
		case pb.Counter != nil:
			values[key] = pb.GetCounter().GetValue()
		case pb.Gauge != nil:
			values[key] = pb.GetGauge().GetValue()
		}
	}
	return values, apiErr
}

func TestFirewallCollectorUpdate(t *testing.T) {
	c, client := newFirewallTestCollector(t, "")

	values, apiErr := collectFirewall(t, c, client)
	if apiErr != nil {
		t.Fatalf("unexpected error: %s", apiErr.Error())
	}

	expected := map[string]float64{
		"opnsense_firewall_in_ipv4_pass_packets|igb0":              10,
		"opnsense_firewall_in_ipv4_pass_bytes_total|igb0":          123456789012,
		"opnsense_firewall_out_ipv6_pass_bytes_total|igb0":         400,
		"opnsense_firewall_interface_references|igb0":              7,
		"opnsense_firewall_state_table_entries":                    512,
		"opnsense_firewall_state_table_operations_total|inserts":   9000,
		"opnsense_firewall_source_tracking_entries":                0,
		"opnsense_firewall_pf_counters_total|match":                42,
		"opnsense_firewall_pf_limit_counters_total|max-src-states": 0,
		"opnsense_firewall_memory_limit|states":                    400000,
		"opnsense_firewall_timeout_seconds|tcp.first":              120,
		"opnsense_firewall_adaptive_states|adaptive.start":         240000,
	}
	for key, want := range expected {
		got, ok := values[key]
		if !ok {
			t.Errorf("missing metric %s", key)
			continue
		}
		if got != want {
			t.Errorf("%s: expected %v, got %v", key, want, got)
		}
	}

	if _, ok := values["opnsense_firewall_interface_counters_cleared_timestamp_seconds|igb0"]; !ok {
		t.Errorf("missing cleared timestamp metric")
	}
}

func TestFirewallCollectorPartialFailure(t *testing.T) {
	c, client := newFirewallTestCollector(t, "/api/diagnostics/firewall/pf_statistics/info")

	values, apiErr := collectFirewall(t, c, client)
	if apiErr == nil {
		t.Fatalf("expected the pf_statistics/info failure to be reported")
	}
	if _, ok := values["opnsense_firewall_pf_counters_total|match"]; ok {
		t.Errorf("did not expect pf counters when pf_statistics/info fails")
	}
	for _, key := range []string{
		"opnsense_firewall_in_ipv4_pass_bytes_total|igb0",
		"opnsense_firewall_memory_limit|states",
		"opnsense_firewall_timeout_seconds|tcp.first",
	} {
		if _, ok := values[key]; !ok {
			t.Errorf("expected %s despite the pf_statistics/info failure", key)
		}
	}
}
