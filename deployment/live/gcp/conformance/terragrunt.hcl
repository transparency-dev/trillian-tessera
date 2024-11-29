terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//conformance"
}

locals {
  env                      = path_relative_to_include()
  project_id               = get_env("GOOGLE_PROJECT")
  location                 = get_env("GOOGLE_REGION", "us-central1")
  base_name                = get_env("TESSERA_BASE_NAME", "${local.env}-conformance")
  server_docker_image      = get_env("TESSERA_CLOUD_RUN_DOCKER_IMAGE")
  signer                   = get_env("TESSERA_SIGNER")
  tessera_reader           = get_env("TESSERA_READER", "")
  tessera_writer           = get_env("TESSERA_WRITER", "")
  conformance_readers      = length(local.tessera_reader) > 0 ? [local.tessera_reader] : []
  conformance_writers      = length(local.tessera_writer) > 0 ? [local.tessera_writer] : []
  cloudrun_service_account = get_env("TESSERA_CLOUD_RUN_SERVICE_ACCOUNT", "")
}

remote_state {
  backend = "gcs"

  config = {
    project  = local.project_id
    location = local.location
    bucket   = "${local.project_id}-${local.base_name}-terraform-state"
    prefix   = "${local.env}/terraform.tfstate"

    gcs_bucket_labels = {
      name = "terraform_state_storage"
    }
  }
}
