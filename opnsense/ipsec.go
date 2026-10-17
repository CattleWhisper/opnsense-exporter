package opnsense

import (
	"encoding/json"
	"strconv"
	"strings"
)

type ipsecSearchResponse struct {
	Rows []struct {
		Phase1desc  string `json:"phase1desc"`
		Connected   bool   `json:"connected"`
		IkeId       string `json:"ikeid"`
		Name        string `json:"name"`
		InstallTime string `json:"install-time"`
		BytesIn     int64  `json:"bytes-in"`
		BytesOut    int64  `json:"bytes-out"`
		PacketsIn   int64  `json:"packets-in"`
		PacketsOut  int64  `json:"packets-out"`
	} `json:"rows"`
	RowCount int `json:"rowCount"`
	Total    int `json:"total"`
	Current  int `json:"current"`
}

type ipsecPhase2 struct {
	Phase2desc  string
	Name        string
	SpiIn       string
	SpiOut      string
	InstallTime int64
	RekeyTime   int64
	LifeTime    int64
	BytesIn     int64
	BytesOut    int64
	PacketsIn   int64
	PacketsOut  int64
}

type ipsecPhase2SearchResponse struct {
	Rows []struct {
		Phase2desc  string `json:"phase2desc"`
		Name        string `json:"name"`
		SpiIn       string `json:"spi-in"`
		SpiOut      string `json:"spi-out"`
		InstallTime string `json:"install-time"`
		RekeyTime   string `json:"rekey-time"`
		LifeTime    string `json:"life-time"`
		BytesIn     string `json:"bytes-in"`
		BytesOut    string `json:"bytes-out"`
		PacketsIn   string `json:"packets-in"`
		PacketsOut  string `json:"packets-out"`
	} `json:"rows"`
}

type IPsec struct {
	Phase1desc  string
	Connected   int
	IkeId       string
	Name        string
	InstallTime int64
	BytesIn     int64
	BytesOut    int64
	PacketsIn   int64
	PacketsOut  int64
	Phase2      []ipsecPhase2
}

type IPsecPhase1 struct {
	Rows []IPsec
}

func (c *Client) FetchIPsecPhase2(ikeId string) (ipsecPhase2SearchResponse, *APICallError) {
	var resp ipsecPhase2SearchResponse

	url, ok := c.endpoints["ipsecPhase2"]

	if !ok {
		return resp, &APICallError{
			Endpoint:   "ipsecPhase2",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	body := map[string]string{"id": ikeId}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return resp, &APICallError{
			Endpoint:   "ipsecPhase2",
			Message:    "failed to marshal body",
			StatusCode: 0,
		}
	}

	if err := c.do("POST", url, strings.NewReader(string(bodyBytes)), &resp); err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Client) FetchIPsecPhase1() (IPsecPhase1, *APICallError) {
	var resp ipsecSearchResponse
	var data IPsecPhase1

	url, ok := c.endpoints["ipsecPhase1"]
	if !ok {
		return data, &APICallError{
			Endpoint:   "ipsecPhase1",
			Message:    "endpoint not found in client endpoints",
			StatusCode: 0,
		}
	}

	if err := c.do("GET", url, nil, &resp); err != nil {
		return data, err
	}

	for _, v := range resp.Rows {

		installTime, err := strconv.ParseInt(v.InstallTime, 10, 64)
		if err != nil {
			installTime = 0
		}

		phase2Rows := []ipsecPhase2{}
		phase2, err2 := c.FetchIPsecPhase2(v.IkeId)
		if err2 != nil {
			c.log.Error("failed to fetch ipsec phase2", "error", err2)
		} else {
			for _, v2 := range phase2.Rows {
				p2InstallTime, err := strconv.ParseInt(v2.InstallTime, 10, 64)
				if err != nil {
					p2InstallTime = 0
				}
				rekeyTime, err := strconv.ParseInt(v2.RekeyTime, 10, 64)
				if err != nil {
					rekeyTime = 0
				}
				lifeTime, err := strconv.ParseInt(v2.LifeTime, 10, 64)
				if err != nil {
					lifeTime = 0
				}
				bytesIn, err := strconv.ParseInt(v2.BytesIn, 10, 64)
				if err != nil {
					bytesIn = 0
				}
				bytesOut, err := strconv.ParseInt(v2.BytesOut, 10, 64)
				if err != nil {
					bytesOut = 0
				}
				packetsIn, err := strconv.ParseInt(v2.PacketsIn, 10, 64)
				if err != nil {
					packetsIn = 0
				}
				packetsOut, err := strconv.ParseInt(v2.PacketsOut, 10, 64)
				if err != nil {
					packetsOut = 0
				}
				phase2Rows = append(phase2Rows, ipsecPhase2{
					Phase2desc:  v2.Phase2desc,
					Name:        v2.Name,
					SpiIn:       v2.SpiIn,
					SpiOut:      v2.SpiOut,
					InstallTime: p2InstallTime,
					RekeyTime:   rekeyTime,
					LifeTime:    lifeTime,
					BytesIn:     bytesIn,
					BytesOut:    bytesOut,
					PacketsIn:   packetsIn,
					PacketsOut:  packetsOut,
				})
			}
		}
		data.Rows = append(data.Rows, IPsec{
			Phase1desc:  v.Phase1desc,
			IkeId:       v.IkeId,
			Name:        v.Name,
			InstallTime: installTime,
			BytesIn:     v.BytesIn,
			BytesOut:    v.BytesOut,
			PacketsIn:   v.PacketsIn,
			PacketsOut:  v.PacketsOut,
			Connected:   parseBoolToInt(v.Connected),
			Phase2:      phase2Rows,
		})
	}

	return data, nil
}
