![hieda-social-card](https://github.com/pinkphantasm/hieda/assets/110753839/b11cfc4d-f2f0-412c-be1c-c9cb1a6249c6)

<div align="center"><strong>Hieda</strong> is an open source solution for electronic document management.</div>

## Requirements

- [Docker](https://www.docker.com/) (`docker-compose`)
- Vendor dependencies
- RSA Certificates

Ensure the vendor dependencies:

```shell
cd src/static_service
./scripts/ensure_vendor.sh
```

Then, follow the instructions provided by the script.

To generate private key (required by the certification service), run the following command:

```shell
openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt > ./src/cert_service/key.pem
```

## Run

App can be run using `docker-compose`:

```shell
docker-compose up
```

or, in some cases:

```shell
docker compose up
```

Now, Hieda is available at [localhost](http://localhost).

## Configuration

You can configure Hieda via `docker-compose.yml`.

Services themselves are configured via environment variables. See `README.md` of each service for more details.
