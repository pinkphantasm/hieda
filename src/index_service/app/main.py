from fastapi import FastAPI

from .routers import views

app = FastAPI()


app.include_router(views.router)
