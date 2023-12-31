import grpc
from  proto import police_pb2_grpc, police_pb2, accident_pb2, accident_pb2_grpc, health_pb2
from google.protobuf import json_format   
import json
from fastapi import FastAPI, HTTPException 
from pydantic import BaseModel
from load_balancer import LoadBalancer
from redis_client import RedisClient
import os 
from datetime import datetime, timedelta
from prometheus_client import Counter, generate_latest


#from dotenv import load_dotenv 

import logging

import threading
from saga_coordonator import SagaTransactionCoordinator

from starlette.responses import Response

 

#load_dotenv()

total_requests = Counter(
    "pad_req_counter",
    "Count the requests made by users.",
    ["Status"]
)

street_hardcoded = "vm 99" 
logger = logging.getLogger("GATEWAY_LOGGER") 
logging.basicConfig(level=logging.INFO) 

Police_LB = LoadBalancer(os.getenv("SERVICE_DISCOVERY_HOST") + ":"  + os.getenv("SERVICE_DISCOVERY_PORT") , "police-reporting",logger)   
Accident_LB = LoadBalancer(os.getenv("SERVICE_DISCOVERY_HOST") + ":"  + os.getenv("SERVICE_DISCOVERY_PORT") , "accident-reporting", logger)


Redis_Client = RedisClient(host = os.getenv("REDIS_HOST"), port=os.getenv("REDIS_PORT")) 


TIMEOUT_SECONDS = 3
REROUTE_THRESHOLD = 2
app = FastAPI() 


@app.get("/metrics")
def get_metrics():
    return Response(content= generate_latest(), media_type="text/plain") 


class UserGeoInfo(BaseModel):
    city: str 
    user_long: float
    user_lat: float
    zoom_index: int

@app.get('/fetchPolice') 
def fetch_police(params: UserGeoInfo,reroute_counter=0):
    if reroute_counter >= REROUTE_THRESHOLD:
        Police_LB.call_service_discovery()
        raise HTTPException(status_code=503,detail="circuit break on reroute") 
    res = list()
    try:
        res = Redis_Client.get_pol_values(params.city,params.user_long, params.user_lat) 
    except Exception as e :
        logger.info(str(e))
    if len(res) != 0:
       return json.loads( json.dumps(res))

    user_info = police_pb2.GetPoliceUserEntry(
        user_long=params.user_long,
        user_lat=params.user_lat,
        zoom_index=params.zoom_index,
        city=params.city
    )

    req = police_pb2.FetchPoliceRequest(user_info=user_info) 
    cl = Police_LB.get_next_stub() 
    if cl is  None:
        raise HTTPException(status_code=503,detail="no service registered!") 
    
    def fun_req(cl,timeout):
        response = cl.FetchPolice(req, timeout=timeout) 
        return json_format.MessageToJson(response)

    try:
        response_json =  fun_req(cl,TIMEOUT_SECONDS)
        try:
            Redis_Client.add_pol_coords(params.city, json.loads(response_json))
        except Exception as e:
            logger.info(str(e))

        total_requests.labels(Status="Success").inc()
        return  json.loads(response_json) 

    except Exception as e:
        if  e is grpc.RpcError and grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
            raise HTTPException(status_code=429, detail=str(e.details()))
        else:
            thread = threading.Thread(
                target=Police_LB.circuit_breakers[(Police_LB.current_index -1) % len(Police_LB.server_stubs)].send_requests,
                args=(fun_req, cl, TIMEOUT_SECONDS)
            )

            thread.start()
            total_requests.labels(Status="Failed").inc()

            return fetch_police(params, reroute_counter + 1 )




class UserGeoInfo(BaseModel):
    city: str
    user_long: float
    user_lat: float
    zoom_index: int


@app.get('/fetchAccidents')
def fetch_accs(params: UserGeoInfo,reroute_counter=0):
    if reroute_counter >= REROUTE_THRESHOLD:
        Accident_LB.call_service_discovery()
        raise HTTPException(status_code=503,detail="circuit break on reroute") 
     
    res = list()
    try:
        res = Redis_Client.get_acc_values(params.city,params.user_long, params.user_lat)
    except Exception as e :
        logger.info(str(e))
    if len(res) != 0:
        return json.loads( json.dumps(res))
    
    user_info = accident_pb2.GetAccidentUserEntry(
        user_long=params.user_long,
        user_lat=params.user_lat,
        zoom_index=params.zoom_index,
        city=params.city
    )

    req = accident_pb2.FetchAccidentRequest(user_info=user_info) 

    cl = Accident_LB.get_next_stub() 
    if cl is  None:
        raise HTTPException(status_code=503,detail="no service registered!") 
    
    def fun_req(cl,timeout):
        response = cl.FetchAccidents(req, timeout=timeout) 
        return json_format.MessageToJson(response)
    try:
        response_json =  fun_req(cl,TIMEOUT_SECONDS)  
        try:
            Redis_Client.add_acc_coords(params.city, json.loads(response_json))
        except Exception as e:
            logger.info(str(e))

        total_requests.labels(Status="Success").inc()
        return  json.loads(response_json)

    except Exception as e:
        if  e is grpc.RpcError and grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
            raise HTTPException(status_code=429, detail=str(e.details()))
        else:
            thread = threading.Thread(
                target=Accident_LB.circuit_breakers[(Accident_LB.current_index -1) % len(Accident_LB.server_stubs)].send_requests,
                args=(fun_req, cl, TIMEOUT_SECONDS)
            )

            thread.start()
            total_requests.labels(Status="Failed").inc()

            return fetch_accs(params, reroute_counter + 1 )





class PolicePostParams(BaseModel):
    pol_long: float
    pol_lat: float 
    city: str

@app.post('/postPolice')
def post_police(params: PolicePostParams,reroute_counter=0):

    if reroute_counter >= REROUTE_THRESHOLD:
        Police_LB.call_service_discovery()
        raise HTTPException(status_code=503,detail="circuit break on reroute") 


    police_info = police_pb2.PostPoliceEntry(
        pol_long=params.pol_long,
        pol_lat=params.pol_lat,
        city=params.city
    )

    req = police_pb2.PostPoliceRequest(police_info=police_info)
    cl = Police_LB.get_next_stub() 
    if cl is  None:
        raise HTTPException(status_code=503,detail="no service registered!")  
    
    def fun_req(cl,timeout):
        response = cl.PostPolice(req, timeout=timeout)
        return  json_format.MessageToJson(response)  

    try:
        response_json =  fun_req(cl,TIMEOUT_SECONDS)  
        try:
            Redis_Client.delete_pol_city_info(params.city)
        except Exception as e:
            logger.info(str(e))   

        total_requests.labels(Status="Success").inc()
        return  json.loads(response_json)
    
    except Exception as e:
        if  e is grpc.RpcError and grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
            raise HTTPException(status_code=429, detail=str(e.details()))
        else:
            thread = threading.Thread(
                target=Police_LB.circuit_breakers[(Police_LB.current_index -1) % len(Police_LB.server_stubs)].send_requests,
                args=(fun_req, cl, TIMEOUT_SECONDS)
            )

            thread.start()
            total_requests.labels(Status="Failed").inc()

            return fetch_police(params, reroute_counter + 1 )



class PostAccidentEntry(BaseModel):
    city: str
    accident_long: float
    accident_lat: float
    cars_involved: int

@app.post('/postAccident')
def post_accident(params: PostAccidentEntry,reroute_counter=0):

    if reroute_counter >= REROUTE_THRESHOLD:
        Accident_LB.call_service_discovery()
        raise HTTPException(status_code=503,detail="circuit break on reroute") 
     

    accident_info = accident_pb2.PostAccidentEntry(
        accident_long=params.accident_long,
        accident_lat=params.accident_lat,
        city=params.city,
        street_name=street_hardcoded,
        cars_involved=params.cars_involved
    )

    req = accident_pb2.PostAccidentRequest(accident_info=accident_info) 
    cl = Accident_LB.get_next_stub() 
    if cl is  None:
        raise HTTPException(status_code=503,detail="no service registered!") 
    
    def fun_req(cl,timeout):
        response = cl.PostAccident(req, timeout=timeout) 
        return json_format.MessageToJson(response)

    try:
        response_json =  fun_req(cl,TIMEOUT_SECONDS) 
        try:
            Redis_Client.delete_acc_city_info(params.city)
        except Exception as e:
            logger.info(str(e))


        total_requests.labels(Status="Success").inc()
        return  json.loads(response_json)

    except Exception as e:
        if  e is grpc.RpcError and grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
            raise HTTPException(status_code=429, detail=str(e.details()))
        else:
            thread = threading.Thread(
                target=Accident_LB.circuit_breakers[(Accident_LB.current_index -1) % len(Accident_LB.server_stubs)].send_requests,
                args=(fun_req, cl, TIMEOUT_SECONDS)
            )

            thread.start()

            total_requests.labels(Status="Failed").inc()

            return post_accident(params, reroute_counter + 1 )






class PoliceConfirmParams(BaseModel):
    city:str
    pol_long: float
    pol_lat: float
    confirmation: bool 

@app.put('/confirmPolice')
def confirm_police(params: PoliceConfirmParams,reroute_counter=0):
    if reroute_counter >= REROUTE_THRESHOLD:
        Police_LB.call_service_discovery()
        raise HTTPException(status_code=503,detail="circuit break on reroute") 
     
    confirm_info = police_pb2.ConfirmPoliceEntry(
        pol_long=params.pol_long,
        pol_lat=params.pol_lat,
        city=params.city,
        confirmation=params.confirmation
    )

    req = police_pb2.ConfirmPoliceRequest(police_info=confirm_info)
    cl = Police_LB.get_next_stub() 
    if cl is  None:
        raise HTTPException(status_code=503,detail="no service registered!") 
    
    def fun_req(cl,timeout):
        response = cl.ConfirmPolice(req, timeout=timeout)
        return  json_format.MessageToJson(response)

    try:
        response_json =  fun_req(cl,TIMEOUT_SECONDS)

        try:
            Redis_Client.delete_pol_city_info(params.city)
        except Exception as e:
            logger.info(str(e))    

        total_requests.labels(Status="Success").inc()
        return  json.loads(response_json)

    except Exception as e:
        if  e is grpc.RpcError and grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
            raise HTTPException(status_code=429, detail=str(e.details()))
        else:
            thread = threading.Thread(
                target=Police_LB.circuit_breakers[(Police_LB.current_index -1) % len(Police_LB.server_stubs)].send_requests,
                args=(fun_req, cl, TIMEOUT_SECONDS)
            )

            thread.start()

            total_requests.labels(Status="Failed").inc()

            return confirm_police(params, reroute_counter + 1 )

        


class ConfirmAccidentEntry(BaseModel):
    city: str
    accident_long: float
    accident_lat: float
    police_confirmation: bool
    accident_confirmation: bool


@app.put('/confirmAccident')
def confirm_accident(params: ConfirmAccidentEntry,reroute_counter=0):
    if reroute_counter >= REROUTE_THRESHOLD:
        Accident_LB.call_service_discovery()
        raise HTTPException(status_code=503,detail="circuit break on reroute") 
     
    confirm_info = accident_pb2.ConfirmAccidentEntry(
        accident_long=params.accident_long,
        accident_lat=params.accident_lat,
        police_confirmation=params.police_confirmation,
        accident_confirmation=params.accident_confirmation
    )

    req = accident_pb2.ConfirmAccidentRequest(info=confirm_info) 

    cl = Accident_LB.get_next_stub() 
    if cl is  None:
        raise HTTPException(status_code=503,detail="no service registered!")
    
    def fun_req(cl,timeout):
        response = cl.ConfirmAccident(req, timeout=timeout) 
        return json_format.MessageToJson(response)

    try:
        response_json =  fun_req(cl,TIMEOUT_SECONDS)    
        try:
            Redis_Client.delete_acc_city_info(params.city)
        except Exception as e:
            logger.info(str(e))    
        total_requests.labels(Status="Success").inc()

        return  json.loads(response_json)

    except Exception as e:
        if  e is grpc.RpcError and grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
            raise HTTPException(status_code=429, detail=str(e.details()))
        else:
            thread = threading.Thread(
                target=Accident_LB.circuit_breakers[(Accident_LB.current_index -1) % len(Accident_LB.server_stubs)].send_requests,
                args=(fun_req, cl, TIMEOUT_SECONDS)
            )

            thread.start()
            total_requests.labels(Status="Failed").inc()
            return fetch_accs(params, reroute_counter + 1 )


@app.get('/status')
def gateway_status():
    overall_status = {}
    overall_status["ready"] = True
    try:
        #
        # police
        #  
        req = health_pb2.HealthRequest()
        try:
            cl = Police_LB.get_next_stub() 
            if cl is None:
                overall_status["police_service"] = "something wrong"
                overall_status["ready"] = False
            else:
                response = cl.HealthCheck(req) 
                response_json = json_format.MessageToJson(response)
                overall_status["police_service"] = json.loads(response_json)
        except Exception as e:
                overall_status["police_service"] =  str(e.code())
                overall_status["ready"] = False
        
        # 
        # accident
        #  

        req = health_pb2.HealthRequest()
        try:
            cl = Accident_LB.get_next_stub() 
            if cl == None:
                overall_status["accident_service"] = "something wrong"
                overall_status["ready"] = False
            else:
                response = cl.HealthCheck(req) 
                response_json = json_format.MessageToJson(response)
                overall_status["accident_service"] = json.loads(response_json)
        except Exception as e:
                overall_status["accident_service"] = str(e.code())
                overall_status["ready"] = False

        return json.loads(json.dumps(overall_status))

    except Exception as e:
            return HTTPException(status_code=500, detail=str(e)) 
    

    
@app.post('/informExternalService') 
def inform_external(params: PostAccidentEntry):

    def wrapped1():
        return post_accident(params=params)

    def wrapped2():
        return fetch_police(params=UserGeoInfo(city=params.city,user_lat=params.accident_lat,user_long=params.accident_long,zoom_index=100))

    def reverse1():
        confirm_accident(params=ConfirmAccidentEntry(city=params.city,accident_lat=params.accident_lat,accident_long=params.accident_long,police_confirmation=True,accident_confirmation=False))
    def reverse2():
        pass

    saga_coordonator = SagaTransactionCoordinator(logger)
    saga_coordonator.add_step(wrapped1, reverse1)
    saga_coordonator.add_step(wrapped2, reverse2)

    try:
        results = saga_coordonator.execute()
        total_requests.labels(Status="Success").inc()
        return  json.loads(json.dumps({ "accident_status" : results[0], "police_reported_in_are" : results[1] }))
    except Exception as e:
        total_requests.labels(Status="Failed").inc()
        raise HTTPException(status_code=500, detail=str(e))




if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.getenv("PORT")))
