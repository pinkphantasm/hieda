package health

type ServiceStatus = string

// Service status.
const (
	StatusOperational ServiceStatus = "Operational" // Status operational.
	StatusDegraded    ServiceStatus = "Degraded"    // Status degraded.
	StatusNotRunning  ServiceStatus = "Not running" // Status not running.
)

// Status response details.
const (
	DetailsOperational = "All systems operational" // Details of the "Operational" status.
	DetailsNotRunning  = "Service is not running"  // Details of the "Not running" status.
	NoDetailsProvided  = "No details provided"     // No details provided by the service.
)
