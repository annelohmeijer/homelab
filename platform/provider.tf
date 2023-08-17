terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.17.0"
    }

    minikube = {
      source  = "scott-the-programmer/minikube"
      version = "0.3.1"
    }
  }
}

provider "minikube" {
  kubernetes_version = "v1.26.3"
}
