variable "vpc_id" {
  default = ""
  type = string
}

variable "security_group" {
  default = ""
  type = string
}

data "aws_vpc" "selected" {
  id = var.vpc_id
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [var.vpc_id]
  }
}

data "aws_subnet" "default" {
  for_each = toset(data.aws_subnets.default.ids)
  id       = each.value
}
