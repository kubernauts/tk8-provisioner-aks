package templates

var VariablesAKS = `
variable "cluster_name" {
  default     = "{{.ClusterName}}"
  type        = "string"
  description = "The name of your AKS Cluster"
}

variable "agent_count" {
  default     = "{{.AgentCount}}"
  type        = "string"
  description = "Number of agents in AKS cluster"
}

variable "location" {
  default = "{{.Location}}"
  type    = "string"

  description = "The Azure Region to deploy AKS"
}

variable "ssh_public_key" {
  default = "{{.SSHKeyPath}}"
  
  description = "SSH public key for accessing agents" 
}

variable "client_id" {
  description = "Service principal appID"
}

variable "client_secret" {
  description = "Service principal password"
}

variable "network_plugin" {
  type = "string"
}

variable "dns_prefix" {
  type = "string"
}

variable "resource_group_name" {
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
`

var VariablesRKE = `
variable "cluster_name" {
  default     = "{{.ClusterName}}"
  type        = "string"
  description = "The name of your EKS Cluster"
}

variable "rke_aws_region" {
  default     = "{{.AWSRegion}}"
  # availabe regions are:
  # us-east-1 (Virginia)
  # us-west-2 (Oregon)
  # eu-west-1 (Irland)
  type        = "string"
  description = "The AWS Region to deploy EKS"
}

variable "rke_node_instance_type" {
  default     = "{{.RKENodeInstanceType}}"
  type        = "string"
  description = "Worker Node EC2 instance type"
}

variable "node_count" {
  default     = {{.NodeCount}}
  type        = "string"
  description = "Autoscaling Desired node capacity"
}

variable "cloud_provider" {
  default = "{{.CloudProvider}}"
  type    = "string"
  description = "cloud provider for rke cluster" 
}

variable "authorization" {
  default = "{{.Authorization}}"
  type    = "string"
  description = "authorization mode in rke cluster"
}

variable "AWS_ACCESS_KEY_ID" {
  description = "AWS Access Key"
}

variable "AWS_SECRET_ACCESS_KEY" {
  description = "AWS Secret Key"
}

variable "AWS_DEFAULT_REGION" {
  description = "AWS Region"
}
`
var VariablesEKS = `
# Variables Configuration
variable "cluster-name" {
  default     = "{{.ClusterName}}"
  type        = "string"
  description = "The name of your EKS Cluster"
}

variable "aws-region" {
  default     = "{{.AWSRegion}}"
  # availabe regions are:
  # us-east-1 (Virginia)
  # us-west-2 (Oregon)
  # eu-west-1 (Irland)
  type        = "string"
  description = "The AWS Region to deploy EKS"
}

variable "node-instance-type" {
  default     = "{{.NodeInstanceType}}"
  type        = "string"
  description = "Worker Node EC2 instance type"
}

variable "desired-capacity" {
  default     = {{.DesiredCapacity}}
  type        = "string"
  description = "Autoscaling Desired node capacity"
}

variable "max-size" {
  default     = {{.AutoScallingMaxSize}}
  type        = "string"
  description = "Autoscaling maximum node capacity"
}

variable "min-size" {
  default     = {{.AutoScallingMinSize}}
  type        = "string"
  description = "Autoscaling Minimum node capacity"
}
variable "public_key_path" {
  default = "{{.KeyPath}}"
  type    = "string"
  description = "path to your eks public key"
}

variable "AWS_ACCESS_KEY_ID" {
  description = "AWS Access Key"
}

variable "AWS_SECRET_ACCESS_KEY" {
  description = "AWS Secret Key"
}

variable "AWS_DEFAULT_REGION" {
  description = "AWS Region"
}
`
var Variables = `
variable "AWS_ACCESS_KEY_ID" {
  description = "AWS Access Key"
}

variable "AWS_SECRET_ACCESS_KEY" {
  description = "AWS Secret Key"
}

variable "AWS_SSH_KEY_NAME" {
  description = "Name of the SSH keypair to use in AWS."
}

variable "AWS_DEFAULT_REGION" {
  description = "AWS Region"
}

//General Cluster Settings

variable "aws_cluster_name" {
  description = "Name of AWS Cluster"
}

data "aws_ami" "distro" {
  most_recent = true

  filter {
    name   = "name"
    values = ["{{.OS}}"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["{{.AmiOwner}}"]
}

//AWS VPC Variables

variable "aws_vpc_cidr_block" {
  description = "CIDR Block for VPC"
}

variable "aws_cidr_subnets_private" {
  description = "CIDR Blocks for private subnets in Availability Zones"
  type = "list"
}

variable "aws_cidr_subnets_public" {
  description = "CIDR Blocks for public subnets in Availability Zones"
  type = "list"
}

//AWS EC2 Settings

variable "aws_bastion_size" {
    description = "EC2 Instance Size of Bastion Host"
}

/*
* AWS EC2 Settings
* The number should be divisable by the number of used
* AWS Availability Zones without an remainder.
*/
variable "aws_kube_master_num" {
    description = "Number of Kubernetes Master Nodes"
}

variable "aws_kube_master_size" {
    description = "Instance size of Kube Master Nodes"
}

variable "aws_etcd_num" {
    description = "Number of etcd Nodes"
}

variable "aws_etcd_size" {
    description = "Instance size of etcd Nodes"
}

variable "aws_kube_worker_num" {
    description = "Number of Kubernetes Worker Nodes"
}

variable "aws_kube_worker_size" {
    description = "Instance size of Kubernetes Worker Nodes"
}

/*
* AWS ELB Settings
*
*/
variable "aws_elb_api_port" {
    description = "Port for AWS ELB"
}

variable "k8s_secure_api_port" {
    description = "Secure Port of K8S API Server"
}

variable "default_tags" {
  description = "Default tags for all resources"
  type = "map"
}
`
