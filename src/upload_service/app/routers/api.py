from fastapi import APIRouter, Depends, UploadFile

from ..internal.s3_adapter import S3Adapter, new_s3_adapter


router = APIRouter()


@router.get("/api/health")
async def health():
    return {"details": "All systems operational"}


@router.post("/api/upload")
async def upload(
    file: UploadFile,
    s3_adapter: S3Adapter = Depends(new_s3_adapter),
):
    file_url = s3_adapter.upload(file.file, file.filename)
    return {"url": file_url}
