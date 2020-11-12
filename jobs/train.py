import fire

from sklearn.datasets import load_breast_cancer
from sklearn.linear_model import RidgeClassifier
from sklearn.metrics import f1_score
from sklearn.model_selection import train_test_split

import pylingo


def run(path="../service/artifacts/breast-cancer-1.h5"):
    dataset = load_breast_cancer()
    x = dataset["data"]
    y = dataset["target"]
    tx, vx, ty, vy = train_test_split(x, y, test_size=0.2)

    model = RidgeClassifier()
    model.fit(tx, ty)
    h = model.predict(vx)

    print(f1_score(vy, h))

    pylingo.dump(model, path)


if __name__ == "__main__":
    fire.Fire(run)
