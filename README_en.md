![phantasmagoria-social-card](https://github.com/pinkphantasm/phantasmagoria/assets/110753839/196c8393-b0ec-4e61-abc0-46fd9666ef97)

**Phantasmagoria** is a set of powerful microservices that perform a variety of tasks: file processing, data conversion, third-party API integration, and more.

## Requirements

- [Docker](https://www.docker.com/) (`docker-compose`)
- Vendor dependencies

Ensure the vendor dependencies:

```shell
cd src/static_service
./scripts/ensure_vendor.sh
```

Then, follow the instructions provided by the script.

## Running

Run using `docker-compose`:

```shell
docker compose up
```

or, in some cases:

```
docker-compose up
```

Now, Phantasmagoria is available at [localhost](http://localhost).

## Configuration

You can configure Phantasmagoria via `docker-compose.yml`.

Services themselves are configured via environment variables. See `README.md` of each service for more details.
