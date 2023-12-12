import threading

class SagaTransactionCoordinator:
    def __init__(self,logger):
        self.logger = logger
        self.steps = []
        self.compensations = []
        self.step_events = []  
        self.error_event = threading.Event()  # error event 
        self.error_index = None  
        self.results = {}  
        self.results_lock = threading.Lock() 

    def add_step(self, step, compensation):
        self.steps.append(step)
        self.compensations.append(compensation)
        self.step_events.append(threading.Event()) 

    def execute_step(self, step_index):
        try:
            result = self.steps[step_index]()
            with self.results_lock:
                self.results[step_index] = result
            self.step_events[step_index].set()
        except Exception as e:
            self.error_event.set()
            self.error_index = step_index
            raise e

    def execute_compensations(self):
        for i in reversed(range(self.error_index)):
            if self.step_events[i].is_set():
                self.compensations[i]()

    def execute(self):
        threads = []
        for i in range(len(self.steps)):
            thread = threading.Thread(target=self.execute_step, args=(i,))
            threads.append(thread)
            thread.start()
            
            if self.error_event.is_set():
                break

        for thread in threads:
            thread.join()

        if self.error_event.is_set():
            self.execute_compensations()
            self.logger.info(f"Step {self.error_index+1} failed, compensations executed")
            raise Exception(f"Step {self.error_index+1} failed, compensations executed")

        return self.results