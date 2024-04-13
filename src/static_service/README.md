# Static service

Static service provides a centralized access to the static files (CSS, JS,
images, etc.) for other services.

Put static files in the `public/` directory. You may want to use
`public/vendor` directory for external dependencies, such as `*.min.js`,
`*.min.css` files, etc.

## Environment

### `ADDR`

**Required**: yes

**Value**: valid HTTP address 
(e.g. `:3000`, `127.0.0.1:8080`)

Address which service is started on.

## API Documentation

If service is running separately: `/docs`.

If service is running behind API Gateway: `/static/docs`. In this case, Swagger
file `/static/api/swagger.gateway.json` should be used.
