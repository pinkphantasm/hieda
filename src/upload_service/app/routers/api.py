from fastapi import APIRouter, Depends, HTTPException, UploadFile, status

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
    try:
        file_url = s3_adapter.upload(file.file, file.filename)
        return {"url": file_url}
    except Exception as e:
        raise HTTPException(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail=f"unexpected error: {e}",
        )
