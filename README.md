_case_

## Introduction

Playground project to learn Kubernetes / Containerization / Cloud DevOps on the fly while building a MLPlatform. 
The frameworks used to deploy the trained model are MLflow, Docker and Kubernetes. 

## `model/`

This folder contains all code required to create an mlflow artifact from existing model (or by retraining). 
To get the model as mlflow artifacts create a virtual environment (e.g. conda) with mlflow in it. Since the model
is already trained, we run the wrap entrypoint, which gets the trained model artifacts and logs them in an mlflow run. 
Thus after creating the virtual env run

```commandline
mlflow run model --entry-point wrap --no-conda
```
this will create a folder `mlruns` in which the run is logged.

### Serve model in container on cluster - working solution 

Serving the model in a container requires building the docker image and deploying to the kubernetes cluster (in this case minikube).
Build the image
```commandline
docker build -t model-image -f kube/serve/Dockerfile --build-arg MODEL_URI=mlruns/0/<run-id>/artifacts/model .
```
Deploy the model as API to the cluster
```commandline
kubectl create -f kube/serve
```
which will deploy a few pods in which the model is served as API, a service that handles the requests,
and an ingress resources that makes them available under `http://$(minikube ip)/model/invocations`. Invoke the model
via `python scripts/invoke.py`

### Not working next steps

Next steps would be an MLflow tracking server running on the cluster to which users can log their experiments, and
an API to which serve requests can be posted. This endpoint would then trigger a pipeline that deploys the model

The tracking server and API can be deployed by building the images with
```commandline
docker build -t api -f kube/api/Dockerfile .
docker build -t tracking-server -f kube/tracking_server/Dockerfile .
```
and deploy them to the cluster with
```commandline
kubectl apply -f kube/api -f kube/tracking_server
```


## To do's for up and running app
1. Run model in container with parameter
2. Create local image registry for docker images
3. Deploy model to pods using local image registry
4. Create Gitlab that runs in pod (or Jenkins)
5. Setup CICD to deploy to cluster
6. Setup MLFLOW tracking
7. Setup ingress
