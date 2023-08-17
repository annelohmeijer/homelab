resource "minikube_cluster" "platform" {
  driver       = "docker"
  cluster_name = "mlplatform"
  addons = [
    "ingress"
  ]
}

resource "kubernetes_deployment" "tracking_server" {
  metadata {
    name = "tracking-server"
    labels = {
      App = "platform"
    }
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        App = "platform"
      }
    }
    template {
      metadata {
        labels = {
          App = "platform"
        }
      }
      spec {
        container {
          image = "vad1mo/hello-world-rest:latest"
          name  = "hello-world"
          port {
            container_port = 5345
          }
          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

