import asyncio

from urllib import parse
from uvicorn import Config, Server

from app.config import config


def parse_addr(addr: str):
    result = parse.urlsplit("//" + addr)
    return result.hostname, result.port


async def main():
    if config.addr is None:
        print("environment variable $ADDR must be specified")
        exit(1)

    host, port = parse_addr(config.addr)

    if host is None or port is None:
        print(f"invalid host:port pair: {host}:{port}")
        exit(1)

    server = Server(
        Config(
            "app.main:app",
            host=host,
            port=port,
            reload=config.development,
        )
    )

    await server.serve()


if __name__ == "__main__":
    asyncio.run(main())
