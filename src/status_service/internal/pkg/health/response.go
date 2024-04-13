package health

// ServiceResponse represents status response from the service.
type ServiceResponse struct {
	Details string `json:"details"` // Details provided by the service.
}

// NewResponse returns a new instance of ServiceResponse.
func NewResponse(details string) *ServiceResponse {
	return &ServiceResponse{
		Details: details,
	}
}
