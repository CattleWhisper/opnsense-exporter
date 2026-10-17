package opnsense

import (
	"strconv"
	"strings"
)

type InterfaceDetails struct {
	Device                    string `json:"device"`
	Driver                    string `json:"driver"`
	Index                     string `json:"index"`
	Flags                     string `json:"flags"`
	PromiscuousListeners      string `json:"promiscuous listeners"`
	SendQueueLength           string `json:"send queue length"`
	SendQueueMaxLength        string `json:"send queue max length"`
	SendQueueDrops            string `json:"send queue drops"`
	Type                      string `json:"type"`
	AddressLength             string `json:"address length"`
	HeaderLength              string `json:"header length"`
	LinkState                 string `json:"link state"`
	Vhid                      string `json:"vhid"`
	Datalen                   string `json:"datalen"`
	MTU                       string `json:"mtu"`
	Metric                    string `json:"metric"`
	LineRate                  string `json:"line rate"`
	PacketsReceived           string `json:"packets received"`
	PacketsTransmitted        string `json:"packets transmitted"`
	BytesReceived             string `json:"bytes received"`
	BytesTransmitted          string `json:"bytes transmitted"`
	OutputErrors              string `json:"output errors"`
	InputErrors               string `json:"input errors"`
	Collisions                string `json:"collisions"`
	MulticastsReceived        string `json:"multicasts received"`
	MulticastsTransmitted     string `json:"multicasts transmitted"`
	InputQueueDrops           string `json:"input queue drops"`
	PacketsForUnknownProtocol string `json:"packets for unknown protocol"`
	HWOffloadCapabilities     string `json:"HW offload capabilities"`
	UptimeAtAttachOrStatReset string `json:"uptime at attach or stat reset"`
	Name                      string `json:"name"`
}

// Interface is the struct returned by the OPNsense API
// when requesting the interfaces. The response is weird json
// that have the interface name as key and the interfaceDetails struct as value
type interfaceResponse struct {
	Interface map[string]InterfaceDetails `json:"interfaces"`
}

type Interface struct {
	Name                  string
	Device                string
	Type                  string
	MTU                   int64
	PacketsReceived       int64
	PacketsTransmitted    int64
	BytesReceived         int64
	BytesTransmitted      int64
	MulticastsReceived    int64
	MulticastsTransmitted int64
	InputErrors           int64
	OutputErrors          int64
	Collisions            int64

	// Optional fields. They are parsed leniently, so a release that
	// omits or reformats one of them does not break the scrape.
	// A nil value means the field was missing or not numeric.
	InputQueueDrops           *float64
	SendQueueDrops            *float64
	SendQueueLength           *float64
	SendQueueMaxLength        *float64
	PacketsForUnknownProtocol *float64
	PromiscuousListeners      *float64
	LinkState                 *float64
	LineRateBits              *float64
	UptimeAtAttachOrStatReset *float64
}

// parseOptionalNumber parses a numeric ifinfo value, dropping an optional
// unit suffix (e.g. "1000000000 bit/s"). It returns nil if the value is
// missing or not a number.
func parseOptionalNumber(value, suffix string) *float64 {
	value = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(value), suffix))
	if value == "" {
		return nil
	}
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil
	}
	return &n
}

type Interfaces struct {
	Interfaces []Interface
}

func (c *Client) FetchInterfaces() (Interfaces, *APICallError) {
	var resp interfaceResponse
	var data Interfaces

	url, ok := c.endpoints["interfaces"]
	if !ok {
		return data, &APICallError{
			Endpoint:   "arp",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	err := c.do("GET", url, nil, &resp)
	if err != nil {
		return data, err
	}

	for _, v := range resp.Interface {

		convertedValues, err := sliceIntToMapStringInt(
			[]string{
				v.MTU, v.BytesReceived, v.BytesTransmitted,
				v.PacketsReceived, v.PacketsTransmitted,
				v.MulticastsReceived, v.MulticastsTransmitted,
				v.InputErrors, v.OutputErrors,
				v.Collisions,
			},
			url,
		)
		if err != nil {
			return data, err
		}

		data.Interfaces = append(data.Interfaces, Interface{
			Name:                  v.Name,
			Device:                v.Device,
			Type:                  v.Type,
			MTU:                   convertedValues[v.MTU],
			BytesReceived:         convertedValues[v.BytesReceived],
			BytesTransmitted:      convertedValues[v.BytesTransmitted],
			PacketsReceived:       convertedValues[v.PacketsReceived],
			PacketsTransmitted:    convertedValues[v.PacketsTransmitted],
			MulticastsReceived:    convertedValues[v.MulticastsReceived],
			MulticastsTransmitted: convertedValues[v.MulticastsTransmitted],
			InputErrors:           convertedValues[v.InputErrors],
			OutputErrors:          convertedValues[v.OutputErrors],
			Collisions:            convertedValues[v.Collisions],

			InputQueueDrops:           parseOptionalNumber(v.InputQueueDrops, ""),
			SendQueueDrops:            parseOptionalNumber(v.SendQueueDrops, ""),
			SendQueueLength:           parseOptionalNumber(v.SendQueueLength, ""),
			SendQueueMaxLength:        parseOptionalNumber(v.SendQueueMaxLength, ""),
			PacketsForUnknownProtocol: parseOptionalNumber(v.PacketsForUnknownProtocol, ""),
			PromiscuousListeners:      parseOptionalNumber(v.PromiscuousListeners, ""),
			LinkState:                 parseOptionalNumber(v.LinkState, ""),
			LineRateBits:              parseOptionalNumber(v.LineRate, "bit/s"),
			UptimeAtAttachOrStatReset: parseOptionalNumber(v.UptimeAtAttachOrStatReset, ""),
		})
	}

	return data, nil
}
