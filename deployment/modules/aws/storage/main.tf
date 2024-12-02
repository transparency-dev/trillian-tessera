locals {
  name = "${var.prefix_name}-${var.base_name}"
}

# Configure the AWS Provider
provider "aws" {
  region = var.region
}

# Resources

## S3 Bucket
resource "aws_s3_bucket" "log_bucket" {
  bucket = "${local.name}-bucket"
  force_destroy = var.ephemeral
}

## Aurora MySQL RDS database
resource "aws_rds_cluster" "log_rds" {
  apply_immediately       = true
  cluster_identifier      = "${local.name}-cluster"
  engine                  = "aurora-mysql"
  # TODO(phboneff): make sure that we want to pin this
  engine_version          = "8.0.mysql_aurora.3.05.2"
  database_name           = "tessera"
  master_username         = "root"
  # TODO(phboneff): move to either random strings / Secret Manager / IAM
  master_password         = "password"
  skip_final_snapshot     = true
  backup_retention_period = 1
}

resource "aws_rds_cluster_instance" "cluster_instances" {
  # TODO(phboneff): make some of these variables and/or
  # tweak some of these.
  count              = 1
  identifier         = "${local.name}-writer-${count.index}"
  cluster_identifier = aws_rds_cluster.log_rds.id
  instance_class     = "db.r5.large"
  engine             = aws_rds_cluster.log_rds.engine
  engine_version     = aws_rds_cluster.log_rds.engine_version
}
