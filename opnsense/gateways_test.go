package opnsense

import (
	"bytes"
	"log/slog"
	"regexp"
	"testing"

	"github.com/AthennaMind/opnsense-exporter/internal/options"
	"github.com/prometheus/common/promslog"
)

func TestParseGatewayStatus(t *testing.T) {
	logger := promslog.NewNopLogger()
	tests := []struct {
		name string
		in   string
		want GatewayStatusType
	}{
		{"online", "Online", GatewayStatusOnline},
		{"offline", "Offline", GatewayStatusOffline},
		{"pending", "Pending", GatewayStatusPeding},
		{"packetloss", "Packetloss", GatewayStatusLoss},
		{"latency", "Latency", GatewayStatusLatency},
		{"forced", "Offline (forced)", GatewayStatusForcedDown},
		{"combined latency packetloss", "Latency, Packetloss", GatewayStatusLoss},
		{"combined online latency", "Online, Latency", GatewayStatusLatency},
		{"unknown", "Mystery", GatewayStatusUnknown},
		{"empty", "", GatewayStatusUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseGatewayStatus(tt.in, logger, tt.in)
			if got != tt.want {
				t.Errorf("parseGatewayStatus(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

// OPNsense reports "~" for delay, stddev and loss while a gateway has no
// measurement; it must map to -1 without a parse warning (#106).
func TestParseGatewayMeasurement(t *testing.T) {
	var logs bytes.Buffer
	client, err := NewClient(options.OPNSenseConfig{
		Protocol: "http", Host: "localhost", APIKey: "k", APISecret: "s",
	}, "test", slog.New(slog.NewTextHandler(&logs, nil)))
	if err != nil {
		t.Fatalf("failed to build client: %v", err)
	}

	tests := []struct {
		value string
		regex *regexp.Regexp
		unit  string
		want  float64
	}{
		{"~", client.gatewayRTTRegex, " ms", -1},
		{" ~ ", client.gatewayLossRegex, " %", -1},
		{"12.345 ms", client.gatewayRTTRegex, " ms", 12.345},
		{"0.0 %", client.gatewayLossRegex, " %", 0},
	}
	for _, tt := range tests {
		if got := client.parseGatewayMeasurement(tt.value, tt.regex, tt.unit, "test"); got != tt.want {
			t.Errorf("parseGatewayMeasurement(%q) = %v, want %v", tt.value, got, tt.want)
		}
	}
	if logs.Len() != 0 {
		t.Errorf("expected no warnings, got: %s", logs.String())
	}
}
