import os

from pydantic import BaseModel


class Config(BaseModel):
    addr: str | None
    development: bool


config = Config(
    addr=os.environ.get("ADDR"),
    development=os.environ.get("MODE") == "DEV",
)
