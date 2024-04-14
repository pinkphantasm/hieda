from fastapi import FastAPI

from .routers import api, index


app = FastAPI()

app.include_router(api.router)
app.include_router(index.router)
