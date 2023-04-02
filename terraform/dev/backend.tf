provider "aws" {
  region  = "eu-west-1"
  profile = "dev-admin"
}

terraform {
  backend "local" {
    path = "terraform.tfstate"
  }
}
