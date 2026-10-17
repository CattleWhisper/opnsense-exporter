package opnsense

import (
	"strconv"
	"strings"
)

type HealthCheckResponse struct {
	System struct {
		Status string `json:"status"`
	} `json:"System"`
	CrashReporter struct {
		Message    string `json:"message"`
		Status     string `json:"status"`
		StatusCode int    `json:"statusCode"`
	} `json:"CrashReporter"`
	Firewall struct {
		Message    string `json:"message"`
		Status     string `json:"status"`
		StatusCode int    `json:"statusCode"`
	} `json:"Firewall"`
	// OPNsense>25.1 has a different structure
	// See https://github.com/AthennaMind/opnsense-exporter/issues/48#issuecomment-2692494735
	Metadata struct {
		System struct {
			Status interface{} `json:"status"`
		} `json:"System"`
		CrashReporter struct {
			Message    string `json:"message"`
			Status     string `json:"status"`
			StatusCode int    `json:"statusCode"`
		} `json:"CrashReporter"`
		Firewall struct {
			Message    string      `json:"message"`
			Status     interface{} `json:"status"`
			StatusCode int         `json:"statusCode"`
		} `json:"Firewall"`
	} `json:"metadata"`
}

// GetMetadataSystemStatus converts the Status field to int, handling both string and int types
func (h *HealthCheckResponse) GetMetadataSystemStatus() int {
	switch v := h.Metadata.System.Status.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case string:
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return 0
}

// GetMetadataFirewallStatus converts the Status field to int, handling both string and int types
func (h *HealthCheckResponse) GetMetadataFirewallStatus() int {
	switch v := h.Metadata.Firewall.Status.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case string:
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return 0
}

const (
	HealthCheckStatusOK = "OK"
	// OPNsense>25.1 has a different value
	// See https://github.com/AthennaMind/opnsense-exporter/issues/48#issuecomment-2692494735
	HealthCheckStatusOK_v25_1 = 2
)

// metadataStatusOK reports whether a metadata status value means OK.
// OPNsense 25.1 reports it as the number 2, while 26.1 went back to the
// string "OK" inside the metadata section.
func metadataStatusOK(status interface{}, numeric int) bool {
	if s, ok := status.(string); ok && strings.EqualFold(strings.TrimSpace(s), HealthCheckStatusOK) {
		return true
	}
	return numeric == HealthCheckStatusOK_v25_1
}

// SystemStatusOK reports whether the health check considers the system
// healthy, in any of the formats OPNsense has used: the legacy top-level
// "OK", the 25.1 numeric metadata status, and the 26.1 string metadata status.
func (h *HealthCheckResponse) SystemStatusOK() bool {
	if h.System.Status == HealthCheckStatusOK {
		return true
	}
	return metadataStatusOK(h.Metadata.System.Status, h.GetMetadataSystemStatus())
}

// FirewallStatusOK reports whether the health check considers the
// firewall healthy. Newer OPNsense releases omit the firewall section
// entirely when there is nothing to report, so an absent section
// counts as healthy instead of producing a false negative.
func (h *HealthCheckResponse) FirewallStatusOK() bool {
	if h.Firewall.Status != "" {
		return h.Firewall.Status == HealthCheckStatusOK
	}
	if h.Metadata.Firewall.Status != nil {
		return metadataStatusOK(h.Metadata.Firewall.Status, h.GetMetadataFirewallStatus())
	}
	return true
}

// HealthCheck checks if the OPNsense is up and running.
func (c *Client) HealthCheck() (HealthCheckResponse, error) {
	var resp HealthCheckResponse

	path, ok := c.endpoints["healthCheck"]

	if !ok {
		return HealthCheckResponse{}, &APICallError{
			Endpoint:   "healthCheck",
			Message:    "endpoint not found",
			StatusCode: 0,
		}
	}

	if err := c.do("GET", path, nil, &resp); err != nil {
		return HealthCheckResponse{}, err
	}

	return resp, nil
}
