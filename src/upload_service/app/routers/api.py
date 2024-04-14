from fastapi import APIRouter, Depends, UploadFile
from fastapi.responses import HTMLResponse

from ..internal.s3_adapter import S3Adapter, new_s3_adapter


router = APIRouter()


@router.get("/api/health", status_code=200)
async def health():
    return {"details": "All systems operationals"}


@router.post("/api/upload", response_class=HTMLResponse)
async def upload(
    file: UploadFile,
    s3_adapter: S3Adapter = Depends(new_s3_adapter),
):
    file_url = s3_adapter.upload(file.file, file.filename)

    return HTMLResponse(
        content=f"""<div class="clip-area">
    <span class="clip-area-text">{file_url}</span>
    <button type="button" class="clip-area-button">Копировать</button>
</div>"""
    )
