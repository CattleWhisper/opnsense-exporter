package opnsense

import (
	"encoding/json"
	"testing"

	"github.com/prometheus/common/promslog"
)

// Several udpbroadcastrelay instances share a name and an empty
// description; they must collapse into one entry (#83).
func TestParseServicesMergesDuplicates(t *testing.T) {
	payload := `{"rows": [
		{"id": "a", "name": "udpbroadcastrelay", "description": "", "running": 1},
		{"id": "b", "name": "udpbroadcastrelay", "description": "", "running": 0},
		{"id": "c", "name": "udpbroadcastrelay", "description": "", "running": 1},
		{"id": "d", "name": "unbound", "description": "Unbound DNS", "running": 1},
		{"id": "e", "name": "udpbroadcastrelay", "description": "mdns", "running": 1}
	]}`

	var resp servicesSearchResponse
	if err := json.Unmarshal([]byte(payload), &resp); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	list, running, stopped := parseServices(resp, promslog.NewNopLogger())

	if running != 4 || stopped != 1 {
		t.Errorf("expected totals 4 running / 1 stopped, got %d / %d", running, stopped)
	}
	if len(list) != 3 {
		t.Fatalf("expected 3 unique services, got %d: %+v", len(list), list)
	}
	if list[0].Name != "udpbroadcastrelay" || list[0].Description != "" || list[0].Status != ServiceStatusStopped {
		t.Errorf("expected the merged relay to report stopped, got %+v", list[0])
	}
	if list[2].Description != "mdns" || list[2].Status != ServiceStatusRunning {
		t.Errorf("expected the described relay to stay separate and running, got %+v", list[2])
	}
}
