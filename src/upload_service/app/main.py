from fastapi import FastAPI

from .routers import api, views


app = FastAPI()

app.include_router(api.router)
app.include_router(views.router)
