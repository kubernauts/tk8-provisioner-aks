# AKS Terraform module

module "aks" {
  source              = "./modules/aks"
  cluster_name        = "${var.cluster_name}"
  location            = "${var.location}"
  agent_count         = "${var.agent_count}"
  ssh_public_key      = "${var.ssh_public_key}"
  client_id           = "${var.client_id}"
  client_secret       = "${var.client_secret}"
  dns_prefix          = "${var.dns_prefix}"
  resource_group_name = "${var.resource_group_name}"
  network_plugin      = "${var.network_plugin}"
  os_disk_size_gb     = "${var.os_disk_size_gb}"
  kubernetes_version  = "${var.kubernetes_version}"
  vm_size             = "${var.vm_size}"
}
