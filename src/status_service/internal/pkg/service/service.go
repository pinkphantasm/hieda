// Package service provides utility struct Service.
package service

// Service represents the service whose status is requested.
type Service struct {
	Addr string // Address of the service.
	Name string // Name of the service.
}
