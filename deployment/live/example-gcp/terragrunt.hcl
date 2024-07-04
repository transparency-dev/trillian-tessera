terraform {
  source = "${get_repo_root()}/deployment/modules/gcs"
}

locals {
  project_id  = "trillian-tessera"
  location      = "us-central1"
  base_name = "example-gcs"
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
