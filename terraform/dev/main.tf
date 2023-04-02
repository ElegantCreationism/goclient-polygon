 terraform {
   backend "local" {
     path = ".terraform/terraform.tfstate"
   }
 }

module "fargate" {
  source = "../modules/fargate"

  app_image          = "goclient-polygon"
  app_name           = "goclient-polygon"
  app_port           = "8080"
  ecs_cluster_name   = ""
  security_group_ids = []
  subnet_ids         = []
  vpc_id             = ""
}
