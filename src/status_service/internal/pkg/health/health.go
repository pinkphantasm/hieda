// Package health provides methods to check status of services.
package health

import "time"

// Health represents status of the service.
type Health struct {
	Addr    string        `json:"addr"`    // Address of the request.
	Name    string        `json:"name"`    // Name of the service.
	Time    time.Time     `json:"time"`    // Time of the status check.
	Status  ServiceStatus `json:"status"`  // Status of the service.
	Details string        `json:"details"` // Detail of the service status.
}

func New(addr, name string, status ServiceStatus, details string) *Health {
	return &Health{
		Addr:    addr,
		Name:    name,
		Time:    time.Now(),
		Status:  status,
		Details: details,
	}
}
