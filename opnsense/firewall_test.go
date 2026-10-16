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
