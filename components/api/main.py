"""Platform backend that handles requests."""
from fastapi import FastAPI

app = FastAPI()

@app.get('/home')
async def home():
    return {"msg": "home"}

@app.post("/invoke")
async def invoke():
    return {"msg": "data"}

