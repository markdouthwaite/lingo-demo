import random
from locust import HttpUser, task
from sklearn.datasets import load_breast_cancer


def iter_dataset():
    dataset = load_breast_cancer()
    x = dataset["data"]
    while True:
        i = random.randint(0, len(x)-1)
        yield x[i].tolist()


class WebsiteUser(HttpUser):
    min_wait = 5000
    max_wait = 9000

    observation = iter_dataset()

    @task(10)
    def invoke(self):
        payload = {"features": next(self.observation)}
        self.client.post("/predict", json=payload)

    @task(1)
    def ping(self):
        self.client.get("/health")
