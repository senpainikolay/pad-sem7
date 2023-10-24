
import requests 
import json 
from main import TIMEOUT_SECONDS
from circuit_breaker import CircuitBreaker
from  proto import police_pb2_grpc, police_pb2, accident_pb2, accident_pb2_grpc, health_pb2 
import grpc

class LoadBalancer:
    def __init__(self, service_discovery_server,service_type):
        self.__reset_states()
        self.service_discovery_server_url = service_discovery_server 
        self.service_type = service_type 

        e = self.call_service_discovery()
        if e != None:
            print(f"SERVICE DISCOVERY DOWN, Exception: {e}") 

    def __reset_states(self):
        self.server_stubs = [] 
        self.server_urls = [] 
        self.circuit_breakers = []
        self.current_index = 0

    def call_service_discovery(self):
        try:
            self.__reset_states()
            res = requests.get("http://" + self.service_discovery_server_url + f"/services/{self.service_type}")  
            data = json.loads(res.text) 
            for s_url in data["urls"]:
                cl = self.register_stub(s_url) 
                if cl is not None:
                    self.server_stubs.append(cl) 
                    self.circuit_breakers.append(CircuitBreaker(reset_timeout=TIMEOUT_SECONDS*3.5))
            return None
        except Exception as e:
            return e
        
    def register_stub(self,url):
        try:
            if self.service_type == "police-reporting":
                cl =  police_pb2_grpc.PoliceReportingServiceStub(grpc.insecure_channel(url)) 
                self.server_urls.append(url)
                return cl
            elif   self.service_type == "accident-reporting":
                cl =  accident_pb2_grpc.AccidentReportingServiceStub(grpc.insecure_channel(url)) 
                self.server_urls.append(url)
                return cl
            return None
        except:
            return None
    
    def unregister_url(self):
          try:
            idx = self.current_index-1
            if idx < 0:
                idx = len(self.server_urls)-1
            data = json.dumps({
                        "service_type" : self.service_type,
                        "service_url" : self.server_urls[idx]
                    })
            headers = {
                "Content-Type": "application/json"
            }
            requests.delete("http://" + self.service_discovery_server_url + f"/services/unregister",data=data,headers=headers) 
          except:
              return None 
          
    def record_to_circuit_breaker(self):
        idx = self.current_index-1
        if idx < 0:
            idx = len(self.server_urls)-1
            
        self.circuit_breakers[idx].record_fail()
        
    def get_next_stub(self):
        if len(self.server_urls) == 0:
            return None
        next_client = self.server_stubs[self.current_index]
        self.current_index = (self.current_index + 1) % len(self.server_stubs)
        return next_client