terraform {
  backend "gcs" {}
}

## Call the Tessera GCP module
##
## This will configure all the storage infrastructure required to run a Tessera log on GCP.
module "gcs" {
  source = "..//gcs"

  base_name  = var.base_name
  env        = var.env
  location   = var.location
  project_id = var.project_id
}

# Enable Cloud Run API
resource "google_project_service" "cloudrun_api" {
  service            = "run.googleapis.com"
  disable_on_destroy = false
}

###
### Set up Cloud Run service
###
resource "google_service_account" "cloudrun_service_account" {
  account_id   = "cloudrun-${var.env}-sa"
  display_name = "Service Account for Cloud Run (${var.env})"
}

resource "google_project_iam_member" "iam_act_as" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.cloudrun_service_account.email}"
}
resource "google_project_iam_member" "iam_metrics_writer" {
  project = var.project_id
  role    = "roles/monitoring.metricWriter"
  member  = "serviceAccount:${google_service_account.cloudrun_service_account.email}"
}
resource "google_spanner_database_iam_binding" "iam_spanner_database_user" {
  project  = var.project_id
  instance = module.gcs.log_spanner_instance.name
  database = module.gcs.log_spanner_db.name
  role     = "roles/spanner.databaseUser"

  members = [
    "serviceAccount:${google_service_account.cloudrun_service_account.email}"
  ]
}
resource "google_project_iam_member" "iam_service_agent" {
  project = var.project_id
  role    = "roles/run.serviceAgent"
  member  = "serviceAccount:${google_service_account.cloudrun_service_account.email}"
}

resource "google_cloud_run_v2_service" "default" {
  name         = "example-service-${var.env}"
  location     = var.location
  launch_stage = "GA"

  template {
    service_account = google_service_account.cloudrun_service_account.email
    containers {
      image = var.example_gcp_docker_image
      name  = "example-gcp"
      args = [
        "--logtostderr",
        "--v=1",
        "--bucket=${module.gcs.log_bucket.id}",
        "--spanner=projects/${var.project_id}/instances/${module.gcs.log_spanner_instance.name}/databases/${module.gcs.log_spanner_db.name}",
        "--project=${var.project_id}",
        "--listen=:8080",
        "--signer=./testgcp.sec",
      ]
      ports {
        container_port = 8080
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
    containers {
      image      = "us-docker.pkg.dev/cloud-ops-agents-artifacts/cloud-run-gmp-sidecar/cloud-run-gmp-sidecar:1.0.0"
      name       = "collector"
      depends_on = ["example-gcp"]
    }
  }
  client = "terraform"
  depends_on = [
    module.gcs,
    google_project_service.cloudrun_api,
    google_project_iam_member.iam_act_as,
    google_project_iam_member.iam_metrics_writer,
    google_project_iam_member.iam_service_agent,
    google_spanner_database_iam_binding.iam_spanner_database_user,
  ]
}

