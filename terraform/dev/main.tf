module "fargate" {
  source = "../modules/fargate"

  app_image          = "elegantcreationism/goclient-polygon:latest"
  app_name           = "goclient-polygon"
  app_port           = "8080"
  ecs_cluster_name   = "DevCluster"
  security_group_ids = [var.security_group]
  subnet_ids         = [for s in data.aws_subnet.default : s.id]
  vpc_id             = data.aws_vpc.selected.id
}
