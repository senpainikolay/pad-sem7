import time

import time

class CircuitBreaker:
    def __init__(self, max_failures=3, reset_timeout=60):
        self.max_failures = max_failures
        self.reset_timeout = reset_timeout
        self.failure_count = 0
        self.last_time = time.time()

    def record_fail(self):
        self.failure_count += 1
        if self.failure_count >= self.max_failures and time.time() - self.last_time >= self.reset_timeout:
            print("CIRCUIT BREAKER TRIP")
            self.__reset()
            return 
        if time.time()  - self.last_time >= self.reset_timeout:
            self.__reset()

    def __reset(self):
        self.failure_count = 0
        self.last_time = time.time()