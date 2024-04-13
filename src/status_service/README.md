# Status service

Status service checks statuses of the specified list of services. It sends
`GET` requests to the `/api/health` of each service.

## Environment

### `ADDR`

**Required**: yes

**Value**: valid HTTP address 
(e.g. `:3000`, `127.0.0.1:8080`)

Address which service is started on.

### `SERVICES`

**Required**: no

**Value**: Space-separated entries `NAME=ADDR`, where `NAME`
is a name of service and `ADDR` is its address
(e.g. `Name=http://127.0.0.1:8080 Another=http://127.0.0.1:3123`)

List of services whose statuses are checked.

## API Documentation

If service is running separately: `/docs`.

If service is running behind API Gateway: `/status/docs`. In this case, Swagger
file `/status/api/swagger.gateway.json` should be used.
