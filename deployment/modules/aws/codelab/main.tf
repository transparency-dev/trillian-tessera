# Header ######################################################################
terraform {
  backend "s3" {}
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.76.0"
    }
  }
}

locals {
  name = "${var.prefix_name}-${var.base_name}"
  port = 2024
}

provider "aws" {
  region = var.region
}

module "storage" {
  source = "../storage"

  prefix_name = var.prefix_name
  base_name   = var.base_name
  region      = var.region
  ephemeral   = true
}

# Resources ####################################################################
## Virtual private network #####################################################
# This will be used for the containers to communicate between themselves, and
# the S3 bucket.
resource "aws_default_vpc" "default" {
   tags = {
    Name = "Default VPC"
  }
}

## Connect S3 bucket to VPC ####################################################
# This allows the hammer to talk to a non public S3 bucket over HTTP.
resource "aws_vpc_endpoint" "s3" {
  vpc_id       = aws_default_vpc.default.id
  service_name = "com.amazonaws.${var.region}.s3"
}

resource "aws_vpc_endpoint_route_table_association" "private_s3" {
  vpc_endpoint_id = aws_vpc_endpoint.s3.id
  route_table_id  = aws_default_vpc.default.default_route_table_id
}

resource "aws_s3_bucket_policy" "allow_access_from_vpce" {
  bucket = module.storage.log_bucket.id
  policy = data.aws_iam_policy_document.allow_access_from_vpce.json
}

data "aws_iam_policy_document" "allow_access_from_vpce" {
  statement {
    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${module.storage.log_bucket.arn}/*",
    ]

    condition {
     test = "StringEquals"
     variable = "aws:sourceVpce" 
     values = [aws_vpc_endpoint.s3.id]
    }
  }
  depends_on = [aws_vpc_endpoint.s3]
}
