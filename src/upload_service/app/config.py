import boto3

import os

from pydantic import BaseModel


class Config(BaseModel):
    addr: str | None
    development: bool
    key: str | None
    secret_key: str | None
    bucket: str | None


config = Config(
    addr=os.environ.get("ADDR"),
    development=os.environ.get("MODE") == "DEV",
    key=os.environ.get("S3_KEY"),
    secret_key=os.environ.get("S3_SECRET_KEY"),
    bucket=os.environ.get("S3_BUCKET"),
)


session = boto3.session.Session()
s3 = session.client(
    service_name="s3",
    endpoint_url="https://storage.yandexcloud.net",
    aws_access_key_id=config.key,
    aws_secret_access_key=config.secret_key,
)
