from fastapi import FastAPI
from matching.model import *
from pydantic import BaseModel

class MatchingRequest(BaseModel):
    student_infos: str
    project_infos: str

class MatchingResponse(BaseModel):
    match_score: float

app = FastAPI()

# API endpoint
# Input json: student_infos, project_infos
# Output json: match_score
@app.post('/matching/', response_model=MatchingResponse)
async def root(request: MatchingRequest) -> any:
    return {
        "match_score": generate_score(request.student_infos, request.project_infos)
    }