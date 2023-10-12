from fastapi import FastAPI, HTTPException 
from pydantic import BaseModel
import json


app = FastAPI()

services = {}

class ServiceInfo(BaseModel):
    service_type: str
    service_url: str


@app.get('/status')
async def gateway_status():
    overall_status = {}
    overall_status["status"] = True
    overall_status["police_service"] = "registered" 
    overall_status["accident_service"] = "registered" 

    if "police-reporting"  not in services:
        overall_status["police_service"] = "unregistered" 
    if "accident-reporting"  not in services:
        overall_status["accident_service"] = "unregistered"

    return json.loads(json.dumps(overall_status))

@app.post("/services/register")
def register_service(params: ServiceInfo):
    if params.service_type not in services:
        services[params.service_type] = list()
    if params.service_url not in services[params.service_type]:
        services[params.service_type].append(params.service_url)
    return {"message": f"Service replica of '{params.service_type}' registered at '{params.service_url}'"}

@app.delete("/services/unregister")
def unregister_service(params: ServiceInfo):
    if params.service_type in services and params.service_url in services[params.service_type]:
        services[params.service_type].remove(params.service_url)
        return {"message": f"Service '{params.service_url}' of type '{params.service_type}' unregistered"}
    else:
        raise HTTPException(status_code=404, detail=f"Service '{params.service_url}' of type '{params.service_type}' not found")

@app.get("/services/{service_type}")
def get_services(service_type: str):
    if service_type in services:
        return {"urls": services[service_type]}
    else:
        raise HTTPException(status_code=404, detail=f"Service type '{service_type}' not found")


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)