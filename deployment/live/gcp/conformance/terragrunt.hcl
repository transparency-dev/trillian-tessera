terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//conformance"
}

locals {
  env                          = path_relative_to_include()
  project_id                   = get_env("GOOGLE_PROJECT", "trillian-tessera")
  location                     = get_env("GOOGLE_REGION", "us-central1")
  base_name                    = get_env("TESSERA_BASE_NAME", "${local.env}-conformance")
  conformance_gcp_docker_image = "${local.location}-docker.pkg.dev/trillian-tessera/docker-${local.env}/conformance-gcp:latest"
  signer                       = get_env("TESSERA_SIGNER")
  verifier                     = get_env("TESSERA_VERIFIER")
  # Service accounts are managed externally:
  conformance_users        = ["serviceAccount:cloudbuild-prod-sa@trillian-tessera.iam.gserviceaccount.com"]
  bucket_readers           = ["serviceAccount:cloudbuild-prod-sa@trillian-tessera.iam.gserviceaccount.com"]
  cloudrun_service_account = "cloudrun-${local.env}-sa@trillian-tessera.iam.gserviceaccount.com"
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
