# Dev environment k3s cluster configuration
include "root" {
  path = find_in_parent_folders()
}

terraform {
  source = "../../../modules/hetzner-k3s"
}

inputs = {
  # Environment-specific overrides
  server_type = "cx21"    # 2 vCPU, 4GB RAM - good for dev
  node_count  = 1         # Single worker for dev to save costs

  # SSH keys - replace with your actual SSH key names from Hetzner Console
  ssh_keys = ["homelab-key"]  # You'll need to create this in Hetzner Console

  # k3s configuration
  k3s_version = "v1.28.2+k3s1"
  enable_traefik = true
  enable_metrics_server = true

  # Networking
  cluster_cidr = "10.42.0.0/16"
  service_cidr = "10.43.0.0/16"

  # This will be passed from environment variables
  # Set via: export TF_VAR_hcloud_token="your_token_here"
  # hcloud_token = # Set via environment variable
}