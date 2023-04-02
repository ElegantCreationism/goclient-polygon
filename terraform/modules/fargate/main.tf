resource "aws_ecs_task_definition" "app_task" {
  family                   = "${var.app_name}-task"
  container_definitions    = <<DEFINITION
[
  {
    "name": "${var.app_name}-container",
    "image": "${var.app_image}",
    "portMappings": [
      {
        "containerPort": ${var.app_port},
        "protocol": "tcp"
      }
    ]
  }
]
DEFINITION
  cpu                      = 256
  memory                   = 512
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
}

resource "aws_ecs_service" "app_service" {
  name            = "${var.app_name}-service"
  cluster         = var.ecs_cluster_name
  task_definition = aws_ecs_task_definition.app_task.arn
  desired_count   = 1

  network_configuration {
    security_groups  = var.security_group_ids
    subnets          = var.subnet_ids
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.app_target_group.arn
    container_name   = "${var.app_name}-container"
    container_port   = var.app_port
  }
}

resource "aws_lb_target_group" "app_target_group" {
  name        = "${var.app_name}-target-group"
  port        = var.app_port
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.vpc_id

  health_check {
    path     = "/"
    protocol = "HTTP"
  }
}
