terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//example-gcp"
}

locals {
  env                      = path_relative_to_include()
  project_id               = get_env("GOOGLE_PROJECT", "trillian-tessera")
  location                 = get_env("GOOGLE_REGION", "us-central1")
  base_name                = get_env("TESSERA_BASE_NAME", "${local.env}-example-gcp")
  example_gcp_docker_image = "us-central1-docker.pkg.dev/trillian-tessera/docker-${local.env}/example-gcp:latest"
  log_origin               = "example-gcp-${local.env}"
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
