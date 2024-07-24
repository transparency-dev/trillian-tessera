terraform {
  source = "${get_repo_root()}/deployment/modules/gcs"
}

locals {
  project_id = get_env("GOOGLE_PROJECT")
  location   = get_env("GOOGLE_REGION", "us-central1")
  base_name   = get_env("TESSERA_BASE_NAME", "tessera-example")
}

inputs = merge(
  local,
  {}
)

remote_state {
  backend = "gcs"

  config = {
    project  = local.project_id
    location = local.location
    bucket   = "${local.project_id}-${local.base_name}-terraform-state"

    gcs_bucket_labels = {
      name  = "terraform_state_storage"
    }
  }
}
