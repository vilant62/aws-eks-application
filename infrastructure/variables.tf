variable "region" {
  description = "AWS region"
  type        = string
  default     = "eu-west-3"
}
variable "cluster_name" {
  type    = string
  default = "eks-project"
}
variable "vpc_cidr" {
  description = "The network address for VPC in CIDR notation"
  type        = string
  default     = "10.10.0.0/16"
}
variable "kubernetes_version" {
  type    = string
  default = "1.33"
}
variable "private_subnets" {
  type    = list(string)
  default = ["10.10.10.0/24", "10.10.20.0/24"]
}
variable "public_subnets" {
  type    = list(string)
  default = ["10.10.0.0/24", "10.10.2.0/24"]
}
