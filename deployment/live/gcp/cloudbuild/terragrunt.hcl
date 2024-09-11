terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//cloudbuild"
}

locals {
  project_id         = "trillian-tessera"
  region             = "us-central1"
  env                = path_relative_to_include()
  kms_key_version_id = get_env("TESSERA_KMS_KEY_VERSION", "projects/${local.project_id}/locations/${local.region}/keyRings/ci-conformance/cryptoKeys/log-signer/cryptoKeyVersions/1")
  log_origin         = "ci-conformance"

}

remote_state {
  backend = "gcs"

  config = {
    project  = local.project_id
    location = local.region
    bucket   = "${local.project_id}-cloudbuild-${local.env}-terraform-state"
    prefix   = "${path_relative_to_include()}-terraform.tfstate"

    gcs_bucket_labels = {
      name = "terraform_state_storage"
    }
  }
}
