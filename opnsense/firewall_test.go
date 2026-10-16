package opnsense

import (
	"encoding/json"
	"testing"
)

func TestFirewallPFStatsResponseUnmarshal(t *testing.T) {
	payload := `{
		"interfaces": {
			"igb0": {
				"references": 12,
				"in4_pass_packets": 100,
				"in4_block_packets": 2,
				"out4_pass_packets": 90,
				"out4_block_packets": 1,
				"in6_pass_packets": 50,
				"in6_block_packets": 3,
				"out6_pass_packets": 40,
				"out6_block_packets": 4,
				"in4_pass_bytes": 123456789012,
				"in4_block_bytes": 1024,
				"out4_pass_bytes": 98765432109,
				"out4_block_bytes": 512,
				"in6_pass_bytes": 5555555555,
				"in6_block_bytes": 256,
				"out6_pass_bytes": 4444444444,
				"out6_block_bytes": 128
			}
		}
	}`

	var resp firewallPFStatsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("failed to unmarshal pf statistics response: %v", err)
	}

	stat, ok := resp.Interface["igb0"]
	if !ok {
		t.Fatalf("expected interface igb0 in response, got %+v", resp.Interface)
	}

	if stat.In4PassPackets != 100 || stat.Out6BlockPackets != 4 {
		t.Errorf("unexpected packet counters: %+v", stat)
	}

	expected := map[string][2]int64{
		"in4_pass_bytes":   {stat.In4PassBytes, 123456789012},
		"in4_block_bytes":  {stat.In4BlockBytes, 1024},
		"out4_pass_bytes":  {stat.Out4PassBytes, 98765432109},
		"out4_block_bytes": {stat.Out4BlockBytes, 512},
		"in6_pass_bytes":   {stat.In6PassBytes, 5555555555},
		"in6_block_bytes":  {stat.In6BlockBytes, 256},
		"out6_pass_bytes":  {stat.Out6PassBytes, 4444444444},
		"out6_block_bytes": {stat.Out6BlockBytes, 128},
	}
	for field, v := range expected {
		if v[0] != v[1] {
			t.Errorf("%s: expected %d, got %d", field, v[1], v[0])
		}
	}
}

func TestParsePFInfo(t *testing.T) {
	payload := `{"info": {
		"uptime": "Enabled for 0 days 01:",
		"state-table": {
			"current-entries": {"total": 1234},
			"searches": {"total": 98765432100, "rate": 1500.2},
			"inserts": {"total": 5000, "rate": 1.1},
			"removals": {"total": 3766, "rate": 0.9}
		},
		"source-tracking-table": {
			"current-entries": {"total": 3}
		},
		"counters": {
			"match": {"total": 777, "rate": 2.0},
			"state-mismatch": {"total": 4, "rate": 0.0}
		},
		"limit-counters": {
			"max-states-per-rule": {"total": 0, "rate": 0.0}
		}
	}}`

	var resp firewallPFInfoResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("failed to unmarshal pf info response: %v", err)
	}
	info := parsePFInfo(resp)

	if info.StateTable["current-entries"] != 1234 || info.StateTable["searches"] != 98765432100 || len(info.StateTable) != 4 {
		t.Errorf("unexpected state table: %+v", info.StateTable)
	}
	if info.SourceTracking["current-entries"] != 3 {
		t.Errorf("unexpected source tracking table: %+v", info.SourceTracking)
	}
	if info.Counters["match"] != 777 || info.Counters["state-mismatch"] != 4 {
		t.Errorf("unexpected counters: %+v", info.Counters)
	}
	if v, ok := info.LimitCounters["max-states-per-rule"]; !ok || v != 0 {
		t.Errorf("unexpected limit counters: %+v", info.LimitCounters)
	}
}

func TestParsePFTimeouts(t *testing.T) {
	payload := `{"timeouts": {
		"tcp.first": "120s",
		"tcp.established": "86400s",
		"interval": "10s",
		"adaptive.start": "60000 states",
		"adaptive.end": "120000 states",
		"weird": "n/a"
	}}`

	var resp firewallPFTimeoutsResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("failed to unmarshal pf timeouts response: %v", err)
	}
	timeouts := parsePFTimeouts(resp)

	if timeouts.Seconds["tcp.first"] != 120 || timeouts.Seconds["tcp.established"] != 86400 || timeouts.Seconds["interval"] != 10 {
		t.Errorf("unexpected timeouts: %+v", timeouts.Seconds)
	}
	if len(timeouts.Seconds) != 3 {
		t.Errorf("expected 3 timeouts in seconds, got %+v", timeouts.Seconds)
	}
	if timeouts.AdaptiveStates["adaptive.start"] != 60000 || timeouts.AdaptiveStates["adaptive.end"] != 120000 || len(timeouts.AdaptiveStates) != 2 {
		t.Errorf("unexpected adaptive thresholds: %+v", timeouts.AdaptiveStates)
	}
}
