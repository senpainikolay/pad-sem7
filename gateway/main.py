import grpc
from  proto import *
from google.protobuf import json_format   
import json
from fastapi import FastAPI, HTTPException  
from pydantic import BaseModel


city_harcoded = "chisinau"

police_channel = grpc.insecure_channel('localhost:6666')
POLICE_SERVICE_STUB = police_pb2_grpc.PoliceReportingServiceStub(police_channel)


app = FastAPI() 

class PoliceRequestParams(BaseModel):
    user_long: float
    user_lat: float
    zoom_index: int

@app.get('/fetchPolice')
async def fetch_police(params: PoliceRequestParams):
    user_info = police_pb2.GetPoliceUserEntry(
        user_long=params.user_long,
        user_lat=params.user_lat,
        zoom_index=params.zoom_index,
        city=city_harcoded
    )

    req = police_pb2.FetchPoliceRequest(user_info=user_info)

    try:
        response = POLICE_SERVICE_STUB.FetchPolice(req) 
        response_json = json_format.MessageToJson(response)
        return json.loads(response_json)
 

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details()))





class PolicePostParams(BaseModel):
    pol_long: float
    pol_lat: float

@app.post('/postPolice')
def post_police(params: PoliceRequestParams):
    police_info = police_pb2.PostPoliceEntry(
        pol_long=params.pol_long,
        pol_lat=params.pol_lat,
        city=city_harcoded
    )

    req = police_pb2.PostPoliceRequest(police_info=police_info)

    try:
        response = POLICE_SERVICE_STUB.PostPolice(req)
        response_json = json_format.MessageToJson(response)  
        return json.loads(response_json)

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details()))




class PoliceConfirmParams(BaseModel):
    pol_long: float
    pol_lat: float
    confirmation: bool 

@app.post('/confirmPolice')
def confirm_police(params: PoliceConfirmParams ):
    confirm_info = police_pb2.ConfirmPoliceEntry(
        pol_long=params.pol_long,
        pol_lat=params.pol_lat,
        city=city_harcoded,
        confirmation=params.confirmation
    )

    req = police_pb2.ConfirmPoliceRequest(police_info=confirm_info)

    try:
        response = POLICE_SERVICE_STUB.ConfirmPolice(req)
        response_json = json_format.MessageToJson(response)  
        return json.loads(response_json)

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details())) 
        

    