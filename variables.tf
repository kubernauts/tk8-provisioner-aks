# Variables Configuration
variable "cluster_name" {
  default     = "tk8-aks"
  type        = "string"
  description = "The name of your AKS Cluster"
}

variable "network_plugin" {
  default = "kubenet"
}

variable "location" {
  type        = "string"
  description = "The Azure Region to deploy AKS"
}

variable "ssh_public_key" {
  default = "~/.ssh/id_rsa.pub"
}

variable "agent_count" {
  default     = 3
  type        = "string"
  description = "Autoscaling Desired node capacity"
}

variable "client_id" {
  description = "Service principal appID"
}

variable "client_secret" {
  description = "Service principal password"
}

variable "resource_group_name" {
  type    = "string"
  default = "aksjco-rg"
}

variable "dns_prefix" {
  type = "string"

  default = "aksjco"
}

variable "os_disk_size_gb" {
  default     = "30"
  type        = "string"
  description = "OS disk size which will be applied to agents"
}

variable "kubernetes_version" {
  default     = "1.12.6"
  type        = "string"
  description = "Kubernetes version to use with AKS"
}

variable "vm_size" {
  default     = "Standard_DS2_v2"
  type        = "string"
  description = "VM size used for creating agents"
}
