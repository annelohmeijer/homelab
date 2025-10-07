# Dev environment kubernetes applications configuration
include "root" {
  path = find_in_parent_folders()
}

terraform {
  source = "../../../modules/kubernetes-apps"
}

# Get cluster outputs from the cluster deployment
dependency "cluster" {
  config_path = "../cluster"
  mock_outputs = {
    cluster_endpoint = "https://mock.example.com:6443"
    kubeconfig_path  = "/tmp/mock-kubeconfig"
  }
}

inputs = {
  # Use outputs from cluster deployment
  cluster_endpoint = dependency.cluster.outputs.cluster_endpoint
  kubeconfig_path  = "/tmp/k3s-kubeconfig"  # You'll need to fetch this from the master node

  # Application configuration
  enable_cert_manager = true
  enable_monitoring   = true

  # Namespace configuration
  monitoring_namespace = "monitoring"
  apps_namespace      = "apps"
}