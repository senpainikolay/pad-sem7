import requests 
import json 
from  proto import police_pb2_grpc, police_pb2, accident_pb2, accident_pb2_grpc, health_pb2 
import grpc
from circuit_breaker import CircuitBreaker

TIMEOUT_SECONDS=3

class LoadBalancer:
    def __init__(self, service_discovery_server,service_type, logger):
        self.service_discovery_server_url = service_discovery_server 
        self.service_type = service_type 
        self.server_stubs = [] 
        self.server_urls = [] 
        self.circuit_breakers = []
        self.current_index = 0
        self.logger = logger 
        self.__call_service_discovery

    def __call_service_discovery(self):
        try:
            self.server_stubs = [] 
            self.server_urls = [] 
            self.current_index = 0
            res = requests.get("http://" + self.service_discovery_server_url + f"/services/{self.service_type}")  
            data = json.loads(res.text) 
            for s_url in data["urls"]:
                cl = self.__register_stub(s_url) 
                if cl is not None:
                    self.server_stubs.append(cl)
                    self.server_urls.append(s_url)
                    self.circuit_breakers.append(CircuitBreaker(3,reset_timeout=60,service_type=self.service_type,logger=self.logger))
        except Exception as e:
            self.logger.info(f"Something wrong with service discovery Exception{e}")

        
    def __register_stub(self,url):
        try:
            if self.service_type == "police-reporting":
                cl =  police_pb2_grpc.PoliceReportingServiceStub(grpc.insecure_channel(url)) 
                try:
                    req = health_pb2.HealthRequest()
                    res = cl.HealthCheck(req, timeout=TIMEOUT_SECONDS)
                    return cl
                except:
                    self.__unregister_url(url)
                    self.logger.info(f"Service  {url} unavailable, removed from service discovery")
                    return None
                    
            elif   self.service_type == "accident-reporting":
                cl =  accident_pb2_grpc.AccidentReportingServiceStub(grpc.insecure_channel(url)) 
                try:
                    req = health_pb2.HealthRequest()
                    res = cl.HealthCheck(req, timeout=TIMEOUT_SECONDS)
                    return cl
                except:
                    self.__unregister_url(url)
                    self.logger.info("Service  {url} unavailable, removed from service discovery")
                    return None
        except Exception as e:
            self.logger.info(f"SERVICE DISCOVERY DOWN, Exception{e}")

    
    def __unregister_url(self,url):
            data = json.dumps({
                        "service_type" : self.service_type,
                        "service_url" : url
                    })
            headers = {
                "Content-Type": "application/json"
            }
            try:
                requests.delete("http://" + self.service_discovery_server_url + f"/services/unregister",data=data,headers=headers) 
            except Exception as e:
                self.logger.info(f"SERVICE DISCOVERY DOWN, Exception{e}")

          
    def get_next_stub(self,c=0):
        if len(self.server_urls) == 0:
            if c == -1:
                self.logger.info("No service Registered!")
                return None
            try:
                self.__call_service_discovery()
                return self.get_next_stub(-1)
            except Exception as e:
                self.logger.info(f"Something wrong with service discovery Exception{e}")
                return None
        next_client = self.server_stubs[self.current_index]
        circuit_state = self.circuit_breakers[self.current_index].get_state()

        self.current_index = (self.current_index + 1) % len(self.server_stubs)


        if circuit_state == "open":
            # stupid check, not to get into an infinite loop based on all Circuit breakers are tripped per service
            if c >= len(self.server_urls):
                return next_client
            return self.get_next_stub(c+1)
        
        return next_client