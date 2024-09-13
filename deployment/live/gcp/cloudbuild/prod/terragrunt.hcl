include "root" {
  path   = find_in_parent_folders()
  expose = true
}

inputs = merge(
  include.root.locals,
  {
    service_account    = "cloudbuild-${var.env}-sa@trillian-tessera.iam.gserviceaccount.com"
    kms_key_version_id = get_env("TESSERA_KMS_KEY_VERSION", "projects/${include.root.locals.project_id}/locations/${include.root.locals.region}/keyRings/ci-conformance/cryptoKeys/log-signer/cryptoKeyVersions/1")
    log_origin         = "ci-conformance"
  }
)
