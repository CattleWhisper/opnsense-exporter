package opnsense

import (
	"encoding/json"
	"testing"
)

func TestFirewallStatusOK(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    bool
	}{
		{
			name:    "legacy format ok",
			payload: `{"System":{"status":"OK"},"Firewall":{"status":"OK","statusCode":2}}`,
			want:    true,
		},
		{
			name:    "legacy format error",
			payload: `{"System":{"status":"OK"},"Firewall":{"status":"Error","statusCode":0,"message":"alert"}}`,
			want:    false,
		},
		{
			name:    "25.1 metadata format ok",
			payload: `{"metadata":{"System":{"status":2},"Firewall":{"status":2}}}`,
			want:    true,
		},
		{
			name:    "25.1 metadata format error",
			payload: `{"metadata":{"System":{"status":2},"Firewall":{"status":0,"message":"alert"}}}`,
			want:    false,
		},
		{
			name:    "26.1 metadata string format ok",
			payload: `{"metadata":{"system":{"status":"OK"},"firewall":{"status":"OK"}}}`,
			want:    true,
		},
		{
			name:    "newer format without firewall section",
			payload: `{"metadata":{"system":{"status":2,"message":"No pending messages","title":"System"},"subsystems":[]}}`,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp HealthCheckResponse
			if err := json.Unmarshal([]byte(tt.payload), &resp); err != nil {
				t.Fatalf("failed to unmarshal payload: %v", err)
			}
			if got := resp.FirewallStatusOK(); got != tt.want {
				t.Errorf("FirewallStatusOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemStatusOK(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    bool
	}{
		{"legacy format ok", `{"System":{"status":"OK"}}`, true},
		{"legacy format error", `{"System":{"status":"Error"}}`, false},
		{"25.1 numeric metadata ok", `{"metadata":{"System":{"status":2}}}`, true},
		{"25.1 numeric metadata error", `{"metadata":{"System":{"status":0}}}`, false},
		// OPNsense 26.1 reports the metadata status as a string again (#119)
		{"26.1 string metadata ok", `{"metadata":{"system":{"status":"OK","message":"No pending messages","title":"System"},"subsystems":[]},"subsystems":[]}`, true},
		{"26.1 string metadata lowercase", `{"metadata":{"system":{"status":"ok"}}}`, true},
		{"26.1 string metadata error", `{"metadata":{"system":{"status":"Error"}}}`, false},
		{"empty response", `{}`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp HealthCheckResponse
			if err := json.Unmarshal([]byte(tt.payload), &resp); err != nil {
				t.Fatalf("failed to unmarshal payload: %v", err)
			}
			if got := resp.SystemStatusOK(); got != tt.want {
				t.Errorf("SystemStatusOK() = %v, want %v", got, tt.want)
			}
		})
	}
}
