# Service template

This is a service template that can be used to add new microservices to an
application.

You need to change the service name in the:

- Go files (check the TODO-comments)
- Dockerfile
- Go module file `go.mod`
- API documentation files `api/swagger.json` and `api/swagger.gateway.json`

## Required endpoints

Each service must implement the `GET /api/health` endpoint to work correctly
with the status microservice. It should return the status `200 OK` if service
is operational, non-200 status code if it is degraded. The response body should
be a JSON of the following structure:

```json
{
    "details": "Detailed information about service status"
}
```

## Documentation

Each service may provide the API documentation (`GET /docs` endpoint).

See [Swagger](https://swagger.io/) files `swagger.json` (for a separate start)
and `swagger.gateway.json` (to run with the API Gateway) as an example.

## Configuration

Each service is configured via environment variables. At a minimum, you need to
specify the address where the service will run.

Check the `internal/pkg/config` module as an example.
