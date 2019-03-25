variable "client_id" {}
variable "client_secret" {}


variable "agent_count" {
    type = "string"
}

variable "ssh_public_key" {
    type = "string"
}

variable "dns_prefix" {
    type = "string"
}

variable cluster_name {
    type = "string"
}

variable "resource_group_name" {
    type = "string"
}

variable "location" {
    type = "string"
}

variable "network_plugin" {
    type = "string"
}

variable "os_disk_size_gb" {
  type = "string"
}

variable "kubernetes_version" {
  type = "string"
}

variable "vm_size" {
  type = "string"
}