terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//conformance"
}

locals {
  env        = path_relative_to_include()
  project_id = get_env("GOOGLE_PROJECT", "trillian-tessera")
  location   = get_env("GOOGLE_REGION", "us-central1")
  base_name  = get_env("TESSERA_BASE_NAME", "${local.env}-conformance")
  log_origin = "conformance-gcp-${local.env}"
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
