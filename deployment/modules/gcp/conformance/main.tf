terraform {
  backend "gcs" {}

  required_providers {
    google = {
      source  = "registry.terraform.io/hashicorp/google"
      version = "6.1.0"
    }
  }
}

## Call the Tessera GCP module
##
## This will configure all the storage infrastructure required to run a Tessera log on GCP.
module "gcs" {
  source = "..//gcs"

  base_name          = var.base_name
  env                = var.env
  location           = var.location
  project_id         = var.project_id
  bucket_readers     = var.bucket_readers
  log_writer_members = ["serviceAccount:${var.cloudrun_service_account}"]
}

##
## Resources
##

# Enable Cloud Run API
resource "google_project_service" "cloudrun_api" {
  service            = "run.googleapis.com"
  disable_on_destroy = false
}
resource "google_project_service" "cloudkms_googleapis_com" {
  service = "cloudkms.googleapis.com"
}

/*
## This KMS config is left here for reference, but commented out to avoid
## attempts to delete and re-create these keys with each of the conformance
## runs.

##
## KMS for log signing
##
resource "google_kms_key_ring" "log_signer" {
  location = var.location
  name     = var.base_name

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_kms_crypto_key" "log_signer" {
  key_ring = google_kms_key_ring.log-signer.id
  name     = "log-signer"
  purpose  = "ASYMMETRIC_SIGN"
  version_template {
    algorithm = "EC_SIGN_ED25519"
  }

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_kms_crypto_key_version" "log_signer" {
  crypto_key = google_kms_crypto_key.log_signer.id

  lifecycle {
    prevent_destroy = true
  }
}
*/


locals {
  spanner_db_full = "projects/${var.project_id}/instances/${module.gcs.log_spanner_instance.name}/databases/${module.gcs.log_spanner_db.name}"
}

resource "google_cloud_run_v2_service" "default" {
  name         = var.base_name
  location     = var.location
  launch_stage = "GA"

  template {
    service_account                  = var.cloudrun_service_account
    max_instance_request_concurrency = 700
    timeout                          = "10s"

    scaling {
      max_instance_count = 3
    }

    containers {
      image = var.server_docker_image
      name  = "conformance"
      args = [
        "--logtostderr",
        "--v=1",
        "--bucket=${module.gcs.log_bucket.id}",
        "--spanner=${local.spanner_db_full}",
        "--project=${var.project_id}",
        "--listen=:8080",
        "--kms_key=${var.kms_key_version_id}",
        "--origin=${var.log_origin}",
      ]
      ports {
        name           = "h2c"
        container_port = 8080
      }

      resources {
        limits = {
          cpu    = "2"
          memory = "1024Mi"
        }
      }

      startup_probe {
        initial_delay_seconds = 1
        timeout_seconds       = 1
        period_seconds        = 10
        failure_threshold     = 3
        tcp_socket {
          port = 8080
        }
      }
    }
  }

  deletion_protection = false

  client = "terraform"
  depends_on = [
    module.gcs,
    google_project_service.cloudrun_api,
  ]
}

resource "google_cloud_run_v2_service_iam_binding" "cloudrun_invoker" {
  location = google_cloud_run_v2_service.default.location
  name     = google_cloud_run_v2_service.default.name
  role     = "roles/run.invoker"
  members  = var.conformance_users
}


