from fastapi import APIRouter, Depends, Request, UploadFile, status
from fastapi.responses import HTMLResponse

from ..internal.s3_adapter import S3Adapter, new_s3_adapter
from ..templates import templates

router = APIRouter()


@router.get("/", response_class=HTMLResponse)
async def index(request: Request):
    return templates.TemplateResponse(
        name="index.html",
        request=request,
    )


@router.post("/", response_class=HTMLResponse)
async def upload(
    file: UploadFile,
    s3_adapter: S3Adapter = Depends(new_s3_adapter),
):
    try:
        file_url = s3_adapter.upload(file.file, file.filename)

        return HTMLResponse(
            content=f"""<div class="clip-area">
    <span class="clip-area-text">{file_url}</span>
    <button type="button" class="clip-area-button">Копировать</button>
</div>"""
        )
    except:
        return HTMLResponse(
            status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
            content="""<div>Что-то пошло не так, повторите попытку позже!</div>""",
        )
