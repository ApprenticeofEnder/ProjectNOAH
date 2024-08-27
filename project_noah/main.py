import uvicorn
from fastapi import FastAPI

from project_noah.db import models
from project_noah.db.database import engine

models.Base.metadata.create_all(bind=engine)

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


uvicorn.run(app, host="0.0.0.0", port=8000)
