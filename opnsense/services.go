package opnsense

import "log/slog"

type servicesSearchResponse struct {
	Rows []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Locked      int    `json:"locked"`
		Running     int    `json:"running"`
	} `json:"rows"`
	Total    int `json:"total"`
	RowCount int `json:"rowCount"`
	Current  int `json:"current"`
}

type ServiceStatus int

const (
	ServiceStatusStopped ServiceStatus = iota
	ServiceStatusRunning
	ServiceStatusUnknown
)

type Service struct {
	Description string
	Name        string
	Status      ServiceStatus
}

type Services struct {
	Services     []Service
	TotalRunning int
	TotalStopped int
}

// parseServices converts the services search response. Several instances
// of one service (e.g. udpbroadcastrelay) can share a name and an empty
// description; emitting each of them would register the same label set
// twice and fail the whole /metrics response. Such instances are merged
// into one entry that reports stopped if any of them is stopped. The
// running/stopped totals still count every instance.
func parseServices(resp servicesSearchResponse, log *slog.Logger) ([]Service, int, int) {
	var list []Service
	var running, stopped int
	index := make(map[[2]string]int)

	for _, service := range resp.Rows {
		switch service.Running {
		case 0:
			stopped++
		case 1:
			running++
		}

		status := ServiceStatus(service.Running)
		key := [2]string{service.Name, service.Description}
		if i, seen := index[key]; seen {
			log.Debug("merging services with the same name and description",
				"component", "opnsense-client", "name", service.Name, "description", service.Description)
			if status == ServiceStatusStopped {
				list[i].Status = ServiceStatusStopped
			}
			continue
		}

		index[key] = len(list)
		list = append(list, Service{
			Status:      status,
			Description: service.Description,
			Name:        service.Name,
		})
	}
	return list, running, stopped
}

func (c *Client) FetchServices() (Services, *APICallError) {
	var resp servicesSearchResponse
	var services Services

	url, ok := c.endpoints["services"]

	if !ok {
		return services, &APICallError{
			Endpoint:   "services",
			Message:    "endpoint not found",
			StatusCode: 0,
		}
	}
	err := c.do("GET", url, nil, &resp)
	if err != nil {
		return services, err
	}
	services.Services, services.TotalRunning, services.TotalStopped = parseServices(resp, c.log)

	return services, nil
}
