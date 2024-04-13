import time

from fastapi import APIRouter, UploadFile
from fastapi.responses import HTMLResponse

from ..config import s3, config


router = APIRouter()


@router.get("/api/health", status_code=200)
async def health():
    return {"details": "All systems operationals"}


@router.post("/api/upload", response_class=HTMLResponse)
async def upload(file: UploadFile):
    filename = f"{int(time.time())}-{file.filename}"
    s3.upload_fileobj(file.file, config.bucket, filename)
    file_url = f"https://storage.yandexcloud.net/hieda/{filename}"

    return HTMLResponse(
        content=f"""<div class="clip-area">
    <span class="clip-area-text">{file_url}</span>
    <button type="button" class="clip-area-button">Копировать</button>
</div>"""
    )
