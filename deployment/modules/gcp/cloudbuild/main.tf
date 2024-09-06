terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.41.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

# This will be configured by terragrunt when deploying
terraform {
  backend "gcs" {}
}

resource "google_artifact_registry_repository" "docker" {
  repository_id = "docker-${var.env}"
  location      = var.region
  description   = "Tessera testing docker images"
  format        = "DOCKER"
}

locals {
  artifact_repo            = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.docker.name}"
  conformance_gcp_docker_image = "${local.artifact_repo}/conformance-gcp"
  hammer_docker_image      = "${local.artifact_repo}/hammer"
}

resource "google_cloudbuild_trigger" "docker" {
  name            = "build-docker-${var.env}"
  service_account = google_service_account.cloudbuild_service_account.id
  location        = var.region

  github {
    owner = "transparency-dev"
    name  = "trillian-tessera"
    push {
      branch = "^main$"
    }
  }

  build {
    step {
      id   = "docker_build_hammer"
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t", "${local.hammer_docker_image}:$SHORT_SHA",
        "-t", "${local.hammer_docker_image}:latest",
        "-f", "./hammer/Dockerfile",
        "."
      ]
      wait_for = ["-"]
    }
    step {
      id   = "docker_build_conformance_gcp"
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t", "${local.conformance_gcp_docker_image}:$SHORT_SHA",
        "-t", "${local.conformance_gcp_docker_image}:latest",
        "-f", "./cmd/conformance/gcp/Dockerfile",
        "."
      ]
      wait_for = ["-"]
    }
    step {
      name = "gcr.io/cloud-builders/docker"
      args = [
        "push",
        "--all-tags",
        local.conformance_gcp_docker_image
      ]
      wait_for = ["docker_build_conformance_gcp"]
    }
    step {
      id   = "terraform_apply_conformance_ci"
      name = "alpine/terragrunt"
      entrypoint = "terragrunt"
      args = [
        "apply",
      ]
      dir = "deployment/live/gcp/conformance/ci"
      env = [
        "TF_IN_AUTOMATION=1",
        "TF_INPUT=false",
        "TF_VAR_project_id=${var.project_id}"
      ]
      wait_for = ["-"]
    }
    options {
      logging = "CLOUD_LOGGING_ONLY"
    }
  }
}

resource "google_service_account" "cloudbuild_service_account" {
  account_id   = "cloudbuild-${var.env}-sa"
  display_name = "Service Account for CloudBuild (${var.env})"
}

resource "google_project_iam_member" "act_as" {
  project = var.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}

resource "google_project_iam_member" "logs_writer" {
  project = var.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}

resource "google_project_iam_member" "artifact_registry_writer" {
  project = var.project_id
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}

resource "google_project_iam_member" "cloudrun_deployer" {
  project = var.project_id
  role    = "roles/run.developer"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}

resource "google_storage_bucket_iam_member" "member" {
  bucket   = "${var.project_id}-cloudbuild-${var.env}-terraform-state"
  role = "roles/storage.objectViewer"
  member = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}
