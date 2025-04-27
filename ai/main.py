from fastapi import FastAPI
from matching.model import *
from pydantic import BaseModel


class MatchingRequest(BaseModel):
    student_infos: str
    project_infos: str


class MatchingResponse(BaseModel):
    match_score: float


class MatchingRequest2(BaseModel):
    student_infos: str
    project_infos: list


class MatchingResponse2(BaseModel):
    match_score: list


app = FastAPI()


# API endpoint
# Input json: student_infos, project_infos
# Output json: match_score
@app.post("/matching/", response_model=MatchingResponse)
async def root(request: MatchingRequest) -> any:
    return {"match_score": generate_score(request.student_infos, request.project_infos)}


@app.post("/matching2/", response_model=MatchingResponse2)
async def root2(request: MatchingRequest2) -> any:
    print("The AI service is called")
    match_score = generate_score_with_many_projects(
        request.student_infos, request.project_infos
    )
    print("Match score: " + match_score)
    return {"match_score": match_score}
