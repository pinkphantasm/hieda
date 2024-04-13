import time

from fastapi import APIRouter
from fastapi import UploadFile

from ..config import s3, config


router = APIRouter()


@router.get("/api/health", status_code=200)
async def health():
    return {"details": "All systems operationals"}


@router.post("/api/upload")
async def upload(file: UploadFile):
    filename = f"{int(time.time())}-{file.filename}"
    s3.upload_fileobj(file.file, config.bucket, filename)
    return f"https://storage.yandexcloud.net/hieda/{filename}"
