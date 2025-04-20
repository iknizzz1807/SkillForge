from fastapi import FastAPI
from model import *
from pydantic import BaseModel

class MatchingRequest(BaseModel):
    student_infos: str
    project_infos: str

app = FastAPI()

# API endpoint
# Input json: student_infos, project_infos
# Output json: match_score
@app.post('/matching/')
async def root(request: MatchingRequest):
    return {
        "match_score": generate_score(request.student_infos, request.project_infos)
    }