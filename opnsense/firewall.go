package opnsense

import (
	"encoding/json"
	"strconv"
	"strings"
)

type FirewallPFStat struct {
	InterfaceName string `json:"interface,omitempty"` // We will populate this field with the key of the map
	References    int    `json:"references"`
	// Cleared is the time the pf counters of this interface were last reset,
	// as an ISO 8601 timestamp without zone in the firewall's local time.
	Cleared string `json:"cleared"`

	In4PassPackets   int64 `json:"in4_pass_packets"`
	In4BlockPackets  int64 `json:"in4_block_packets"`
	Out4PassPackets  int64 `json:"out4_pass_packets"`
	Out4BlockPackets int64 `json:"out4_block_packets"`

	In6PassPackets   int64 `json:"in6_pass_packets"`
	In6BlockPackets  int64 `json:"in6_block_packets"`
	Out6PassPackets  int64 `json:"out6_pass_packets"`
	Out6BlockPackets int64 `json:"out6_block_packets"`

	In4PassBytes   int64 `json:"in4_pass_bytes"`
	In4BlockBytes  int64 `json:"in4_block_bytes"`
	Out4PassBytes  int64 `json:"out4_pass_bytes"`
	Out4BlockBytes int64 `json:"out4_block_bytes"`

	In6PassBytes   int64 `json:"in6_pass_bytes"`
	In6BlockBytes  int64 `json:"in6_block_bytes"`
	Out6PassBytes  int64 `json:"out6_pass_bytes"`
	Out6BlockBytes int64 `json:"out6_block_bytes"`
}

// firewallPFStatsResponse is the struct returned by the OPNsense API
// when requesting the firewwall statistics by interface. The response is weird json
// that have the interface name as key and the FirewallPFStats struct as value
type firewallPFStatsResponse struct {
	Interface map[string]FirewallPFStat `json:"interfaces"`
}

type FirewallPFStats struct {
	Interfaces []FirewallPFStat
}

func (c *Client) FetchPFStatsByInterface() (FirewallPFStats, *APICallError) {
	var resp firewallPFStatsResponse
	var data FirewallPFStats

	url, ok := c.endpoints["pfStatisticsByInterface"]
	if !ok {
		return data, &APICallError{
			Endpoint:   "pfStatisticsByInterface",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	err := c.do("GET", url, nil, &resp)
	if err != nil {
		return data, err
	}

	for k, v := range resp.Interface {
		v.InterfaceName = k
		data.Interfaces = append(data.Interfaces, v)
	}
	return data, nil
}

// FirewallPFInfo holds the global pf counters from `pfctl -vvsinfo`,
// keyed by counter name (e.g. "current-entries", "inserts", "match").
type FirewallPFInfo struct {
	StateTable     map[string]float64
	SourceTracking map[string]float64
	Counters       map[string]float64
	LimitCounters  map[string]float64
}

type pfInfoCounter struct {
	Total float64 `json:"total"`
}

type firewallPFInfoResponse struct {
	Info map[string]json.RawMessage `json:"info"`
}

func parsePFInfo(resp firewallPFInfoResponse) FirewallPFInfo {
	data := FirewallPFInfo{
		StateTable:     map[string]float64{},
		SourceTracking: map[string]float64{},
		Counters:       map[string]float64{},
		LimitCounters:  map[string]float64{},
	}
	sections := map[string]map[string]float64{
		"state-table":           data.StateTable,
		"source-tracking-table": data.SourceTracking,
		"counters":              data.Counters,
		"limit-counters":        data.LimitCounters,
	}

	for name, target := range sections {
		raw, ok := resp.Info[name]
		if !ok {
			continue
		}
		var counters map[string]pfInfoCounter
		if err := json.Unmarshal(raw, &counters); err != nil {
			continue
		}
		for k, v := range counters {
			target[k] = v.Total
		}
	}
	return data
}

// FetchPFInfo returns the global pf counters.
func (c *Client) FetchPFInfo() (FirewallPFInfo, *APICallError) {
	var resp firewallPFInfoResponse

	url, ok := c.endpoints["pfStatisticsInfo"]
	if !ok {
		return FirewallPFInfo{}, &APICallError{
			Endpoint:   "pfStatisticsInfo",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	if err := c.do("GET", url, nil, &resp); err != nil {
		return FirewallPFInfo{}, err
	}

	return parsePFInfo(resp), nil
}

// FirewallPFLimits holds the pf memory pool hard limits from `pfctl -vvsmemory`,
// keyed by pool name (e.g. "states", "src-nodes", "frags", "table-entries").
type FirewallPFLimits struct {
	Limits map[string]int64
}

type firewallPFMemoryResponse struct {
	Memory map[string]int64 `json:"memory"`
}

// FetchPFMemoryLimits returns the pf memory pool hard limits.
func (c *Client) FetchPFMemoryLimits() (FirewallPFLimits, *APICallError) {
	var resp firewallPFMemoryResponse

	url, ok := c.endpoints["pfStatisticsMemory"]
	if !ok {
		return FirewallPFLimits{}, &APICallError{
			Endpoint:   "pfStatisticsMemory",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	if err := c.do("GET", url, nil, &resp); err != nil {
		return FirewallPFLimits{}, err
	}

	return FirewallPFLimits{Limits: resp.Memory}, nil
}

// FirewallPFTimeouts holds the pf timeouts from `pfctl -vvstimeouts`.
// Values such as "tcp.first 120s" go to Seconds, and the adaptive
// thresholds ("adaptive.start 60000 states") go to AdaptiveStates.
type FirewallPFTimeouts struct {
	Seconds        map[string]float64
	AdaptiveStates map[string]float64
}

type firewallPFTimeoutsResponse struct {
	Timeouts map[string]string `json:"timeouts"`
}

func parsePFTimeouts(resp firewallPFTimeoutsResponse) FirewallPFTimeouts {
	data := FirewallPFTimeouts{
		Seconds:        map[string]float64{},
		AdaptiveStates: map[string]float64{},
	}
	for name, raw := range resp.Timeouts {
		value := strings.TrimSpace(raw)
		if v, ok := strings.CutSuffix(value, "states"); ok {
			if n, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
				data.AdaptiveStates[name] = n
			}
			continue
		}
		if v, ok := strings.CutSuffix(value, "s"); ok {
			if n, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
				data.Seconds[name] = n
			}
		}
	}
	return data
}

// FetchPFTimeouts returns the pf timeouts.
func (c *Client) FetchPFTimeouts() (FirewallPFTimeouts, *APICallError) {
	var resp firewallPFTimeoutsResponse

	url, ok := c.endpoints["pfStatisticsTimeouts"]
	if !ok {
		return FirewallPFTimeouts{}, &APICallError{
			Endpoint:   "pfStatisticsTimeouts",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	if err := c.do("GET", url, nil, &resp); err != nil {
		return FirewallPFTimeouts{}, err
	}

	return parsePFTimeouts(resp), nil
}
