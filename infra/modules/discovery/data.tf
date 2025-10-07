# Check existing servers
data "hcloud_servers" "existing" {}

# Or by name if you know it
data "hcloud_server" "my_server" {
  name = "your-server-name"
}

# Check existing networks, volumes, etc.
data "hcloud_networks" "existing" {}
data "hcloud_volumes" "existing" {}
