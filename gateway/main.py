import grpc
from  proto import police_pb2_grpc, police_pb2, accident_pb2, accident_pb2_grpc, health_pb2
from google.protobuf import json_format   
import json
from fastapi import FastAPI, HTTPException  
from pydantic import BaseModel


city_harcoded = "chisinau"
street_hardcoded = "vm 99"

TIMEOUT_SECONDS = 5

police_channel = grpc.insecure_channel('localhost:6666')
POLICE_SERVICE_STUB = police_pb2_grpc.PoliceReportingServiceStub(police_channel) 

accident_channel = grpc.insecure_channel('localhost:6667')
ACCIDENT_SERVICE_STUB = accident_pb2_grpc.AccidentReportingServiceStub(accident_channel)


app = FastAPI() 


@app.get('/status')
async def gateway_status():
    overall_status = {}
    overall_status["ready"] = True
    try:
        #
        # police
        #  
        req = health_pb2.HealthRequest()
        try:
            response = POLICE_SERVICE_STUB.HealthCheck(req) 
            print(response)
            response_json = json_format.MessageToJson(response)
            overall_status["police_service"] = json.loads(response_json)
        except grpc.RpcError as e:
                overall_status["police_service"] =  json.loads(json.dumps( {"msg" : e.details(), "error" : True}))
                overall_status["ready"] = False
        
        # 
        # accident
        #  

        req = health_pb2.HealthRequest()
        try:
            response = ACCIDENT_SERVICE_STUB.HealthCheck(req) 
            response_json = json_format.MessageToJson(response)
            overall_status["accident_service"] = json.loads(response_json)
        except grpc.RpcError as e:
                overall_status["accident_service"] = json.loads(json.dumps( {"msg" : e.details(), "error" : True}))
                overall_status["ready"] = False
        
        return json.loads(json.dumps(overall_status))

    except Exception as e:
            return HTTPException(status_code=500, detail=str(e))


class UserGeoInfo(BaseModel):
    user_long: float
    user_lat: float
    zoom_index: int

@app.get('/fetchPolice')
async def fetch_police(params: UserGeoInfo):
    user_info = police_pb2.GetPoliceUserEntry(
        user_long=params.user_long,
        user_lat=params.user_lat,
        zoom_index=params.zoom_index,
        city=city_harcoded
    )

    req = police_pb2.FetchPoliceRequest(user_info=user_info)

    try:
        response = POLICE_SERVICE_STUB.FetchPolice(req, timeout=TIMEOUT_SECONDS) 
        response_json = json_format.MessageToJson(response)
        return  json.loads(response_json)

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details()))



class UserGeoInfo(BaseModel):
    user_long: float
    user_lat: float
    zoom_index: int


@app.get('/fetchAccidents')
async def fetch_police(params: UserGeoInfo):
    user_info = accident_pb2.GetAccidentUserEntry(
        user_long=params.user_long,
        user_lat=params.user_lat,
        zoom_index=params.zoom_index,
        city=city_harcoded
    )

    req = accident_pb2.FetchAccidentRequest(user_info=user_info)

    try:
        response = ACCIDENT_SERVICE_STUB.FetchAccidents(req, timeout=TIMEOUT_SECONDS) 
        response_json = json_format.MessageToJson(response)
        return  json.loads(response_json)
 

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details()))





class PolicePostParams(BaseModel):
    pol_long: float
    pol_lat: float 

@app.post('/postPolice')
def post_police(params: PolicePostParams):
    police_info = police_pb2.PostPoliceEntry(
        pol_long=params.pol_long,
        pol_lat=params.pol_lat,
        city=city_harcoded
    )

    req = police_pb2.PostPoliceRequest(police_info=police_info)

    try:
        response = POLICE_SERVICE_STUB.PostPolice(req, timeout=TIMEOUT_SECONDS)
        response_json = json_format.MessageToJson(response)  
        return  json.loads(response_json)

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details()))



class PostAccidentEntry(BaseModel):
    accident_long: float
    accident_lat: float
    cars_involved: int

@app.post('/postAccident')
async def post_accident(params: PostAccidentEntry):
    accident_info = accident_pb2.PostAccidentEntry(
        accident_long=params.accident_long,
        accident_lat=params.accident_lat,
        city=city_harcoded,
        street_name=street_hardcoded,
        cars_involved=params.cars_involved
    )

    req = accident_pb2.PostAccidentRequest(accident_info=accident_info)

    try:
        response = ACCIDENT_SERVICE_STUB.PostAccident(req, timeout=TIMEOUT_SECONDS)
        response_json = json_format.MessageToJson(response)
        return  json.loads(response_json)

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
        response = POLICE_SERVICE_STUB.ConfirmPolice(req, timeout=TIMEOUT_SECONDS)
        response_json = json_format.MessageToJson(response)  
        return  json.loads(response_json)

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details())) 
        


class ConfirmAccidentEntry(BaseModel):
    accident_long: float
    accident_lat: float
    police_confirmation: bool
    accident_confirmation: bool



@app.post('/confirmAccident')
async def confirm_accident(entry: ConfirmAccidentEntry):
    confirm_info = accident_pb2.ConfirmAccidentEntry(
        accident_long=entry.accident_long,
        accident_lat=entry.accident_lat,
        police_confirmation=entry.police_confirmation,
        accident_confirmation=entry.accident_confirmation
    )

    req = accident_pb2.ConfirmAccidentRequest(info=confirm_info)

    try:
        response = ACCIDENT_SERVICE_STUB.ConfirmAccident(req, timeout=TIMEOUT_SECONDS)
        response_json = json_format.MessageToJson(response)
        return  json.loads(response_json)

    except grpc.RpcError as e:
        if "rate limit exceeded" in str(e.details()):
            raise HTTPException(status_code=429, detail="Rate limit exceeded")
        else:
            raise HTTPException(status_code=500, detail=str(e.details())) 


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8080)