# Microservices Template

A simple and flexible template of the microservice architecture, using:

- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [Fiber](https://github.com/gofiber/fiber)
- [Traefik](https://github.com/traefik/traefik) API Gateway

In this template, services communicate via network requests, using JSON as the
data transfer format.

## Run

App can be run using `docker-compose`:

```shell
docker compose up
```

## Structure

The services are located in the `src/` directory. Services themselves follows
the [Go Project Layout](https://github.com/golang-standards/project-layout/)
specification.

Each service can run separately, if neccessary environment variables are specified.

Check the service template (`src/service_template`) for more information about
services.

## Configuration

The application and its services are configured via environment variables.
See each service `README.md` file for more details.
