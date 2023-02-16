_case_

## Introduction

Playground project to learn Kubernetes, Helm, Terraform, DevOps, on the fly while building a MLPlatform. There are few services to deploy via kube, helm or tf:
- model served as API
- mlflow tracking server 
- API to interact with (e.g. Flask or Django)

These all run on the cluster.

## Setup Ingress and opening a tunnel

In order to make the resources on the cluster available you need do create an ingress instance that redirects
the traffic to the appropriate ports
```buildoutcfg
minikube addons enable ingress
kubectl apply -f kube/ingress.yaml
```
When on Mac with Docker Desktop, you need to setup a tunnel to open up the traffic, since Docker on Mac ...
```buildoutcfg
minikube tunnel
```
Then the services are available at
- http://localhost/tracking-server/#
- http://localhost/api/

## To do's for up and running app
1. Run model in container with parameter
2. Create local image registry for docker images
3. Deploy model to pods using local image registry
4. Create Gitlab that runs in pod (or Jenkins)
5. Setup CICD to deploy to cluster
6. Setup MLFLOW tracking
7. Setup ingress
