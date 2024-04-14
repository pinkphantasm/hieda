from uuid import uuid4

from ..config import s3, config


class S3Adapter:
    def __init__(self, s3, bucket, download_url) -> None:
        self.s3 = s3
        self.bucket = bucket
        self.download_url = download_url

    def upload(self, file, filename):
        object_key = f"hieda_{uuid4()}_{filename}"
        self.s3.upload_fileobj(file, self.bucket, object_key)
        return f"{self.download_url}/{object_key}"


def new_s3_adapter():
    return S3Adapter(s3, config.bucket, config.download_url)
