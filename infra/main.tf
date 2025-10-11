resource "hcloud_ssh_key" "default" {
  name       = "k3s-key"
  public_key = file("~/.ssh/homelab.pub")
}

resource "hcloud_network" "homelab" {
  name = "homelab-network"
  # this provides 65K addresses (2 ^ 16 - 2)
  ip_range = "10.0.0.0/16"

}
resource "hcloud_server" "homelab" {
  name        = "homelab"
  server_type = "cx22"
  image       = "ubuntu-22.04"
  location    = "nbg1" # Nuremberg
  ssh_keys    = [hcloud_ssh_key.default.id]

  labels = {
    role = "server"
  }
}

output "server_public_ip" {
  value = hcloud_server.homelab.ipv4_address

}
