variable "app_name" {
  description = "Name of the application"
}

variable "app_image" {
  description = "Docker image for the application"
}

variable "app_port" {
  description = "Port to expose in the container"
}

variable "vpc_id" {
  description = "ID of the VPC where the service will be deployed"
}

variable "subnet_ids" {
  type        = list(string)
  description = "IDs of the subnets where the service will be deployed"
}

variable "security_group_ids" {
  type        = list(string)
  description = "IDs of the security groups to associate with the service"
}

variable "ecs_cluster_name" {
  description = "Name of the ECS cluster where the service will be deployed"
}
