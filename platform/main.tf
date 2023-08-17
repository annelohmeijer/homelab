resource "minikube_cluster" "platform" {
  driver       = "docker"
  cluster_name = "mlplatform"
  addons = [
    "ingress"
  ]
}
