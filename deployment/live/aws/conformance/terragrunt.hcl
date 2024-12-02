terraform {
  source = "${get_repo_root()}/deployment/modules/aws//conformance"
}

locals {
  env                        = path_relative_to_include()
  account_id                 = "${get_aws_account_id()}"
  region                     = get_env("AWS_REGION", "us-east-1")
  base_name                  = get_env("TESSERA_BASE_NAME", "${local.env}-conformance")
  prefix_name                = get_env("TESSERA_PREFIX_NAME", "trillian-tessera")
  ecr_registry               = get_env("ECR_REGISTRY", "${local.env}.dkr.ecr.${local.region}.amazonaws.com")
  ecr_repository_conformance = get_env("ECR_REPOSITORY_CONFORMANCE", "trillian-tessera/conformance:latest")
  ecr_repository_hammer      = get_env("ECR_REPOSITORY_HAMMER", "trillian-tessera/hammer:latest")
  signer                     = get_env("TESSERA_SIGNER")
  verifier                   = get_env("TESSERA_VERIFIER")
  # Roles are defined externally
  ecs_role  = "arn:aws:iam::${local.account_id}:role/ecsTaskExecutionRole"
  ephemeral = true
}

remote_state {
  backend = "s3"

  config = {
    region         = local.region
    bucket         = "${local.prefix_name}-${local.base_name}-terraform-state"
    key            = "${local.env}/terraform.tfstate"
    dynamodb_table = "${local.prefix_name}-${local.base_name}-terraform-lock"
    s3_bucket_tags = {
      name = "terraform_state_storage"
    }
  }
}
