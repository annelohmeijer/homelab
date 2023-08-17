_case_

## ML Platform on Minikube

Run scalable ML Platform on a local Kubernetes cluster (Minikube) with Terraform. 

### Prerequisites


Playground project to learn Kubernetes, Helm, Terraform, DevOps, on the fly while building a MLPlatform. There are few services to deploy via kube, helm or tf:
- model served as API
- mlflow tracking server 
- API to interact with (e.g. Flask or Django)

These all run on the cluster.

### Components

Run 
```terraform init && terraform plan && terraform apply```
to spin up a Minikube cluster with an ML tracking server as one of the components, which is served an accessible in your browser.


