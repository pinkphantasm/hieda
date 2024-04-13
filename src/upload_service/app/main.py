import time

from fastapi import FastAPI, UploadFile
from .config import s3, config

app = FastAPI()


@app.get('/api/health', status_code=200)
async def health():
    return {'details': 'All systems operationals'}


@app.post('/api/upload')
async def upload(file: UploadFile):
    filename = f'{int(time.time())}-{file.filename}'
    s3.upload_fileobj(file.file, config.bucket, filename)
    return f"https://storage.yandexcloud.net/hieda/{filename}"
