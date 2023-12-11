import grpc
import time

class CircuitBreaker:
    def __init__(self, max_failures, reset_timeout,service_type,logger):
        self.max_failures = max_failures
        self.last_failed = time.time()  -  reset_timeout
        self.reset_timeout = reset_timeout
        self.failures = 0
        self.nr_reqs = 10
        self.state = "closed"
        self.logger = logger
        self.service_type = service_type
    
    def get_state(self):
        if  time.time() - self.last_failed > self.reset_timeout:
            self.state = "closed"
            self.failures = 0
            self.last_failed =  time.time()  -  self.reset_timeout
        return self.state


    def handle_failure(self):
        self.failures += 1
        if self.failures >= self.max_failures:
            self.trip()
            return True
        return False

    def trip(self):
        self.state = "open"
        self.last_failed = time.time()
        self.logger.info(f"Circuit breaker tripped on {self.service_type} service instance")

    def send_requests(self, func_with_client, *args):
        for _ in range(self.nr_reqs):
            try:
                res = func_with_client(*args)
            except grpc.RpcError as e:
                if  grpc.StatusCode.RESOURCE_EXHAUSTED == e.code():
                     continue
                else:
                    failed = self.handle_failure() 
                    if failed:
                        break