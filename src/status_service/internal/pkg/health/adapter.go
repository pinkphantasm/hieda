package health

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pinkphantasm/hieda/src/status_service/internal/pkg/service"
)

// Adapter is an adapter for service status requests.
type Adapter struct {
	c *http.Client // Underlying HTTP client.
}

// NewAdapter returns a new instance of Adapter.
func NewAdapter() *Adapter {
	return &Adapter{
		c: http.DefaultClient,
	}
}

// NewAdapterWithClient returns a new instance of Adapter with underlying http.Client c.
func NewAdapterWithClient(c *http.Client) *Adapter {
	return &Adapter{
		c: c,
	}
}

// CheckHealth checks health of specified services.
// It makes requests to each service's `GET /api/health` endpoint.
func (a Adapter) CheckHealth(services []*service.Service) []*Health {
	var responses []*Health

	for _, srv := range services {
		apiUrl, err := url.JoinPath(srv.Addr, "api", "health")
		if err != nil {
			continue
		}

		r, _ := http.NewRequest(http.MethodGet, apiUrl, nil)

		res, err := a.c.Do(r)

		srvHealth := New(
			srv.Addr,
			srv.Name,
			StatusOperational,
			DetailsOperational,
		)

		if err != nil {
			// Service is not running
			srvHealth.Status = StatusNotRunning
			srvHealth.Details = err.Error()
			responses = append(responses, srvHealth)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			// Service degraded (non-200 status returned)
			srvHealth.Status = StatusDegraded
			srvHealth.Details = NoDetailsProvided
		}

		healthResponse := &ServiceResponse{}
		if err := json.NewDecoder(res.Body).Decode(healthResponse); err == nil {
			srvHealth.Details = healthResponse.Details
		}

		responses = append(responses, srvHealth)
	}

	return responses
}
