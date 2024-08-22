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
  description   = "Tessera example docker images"
  format        = "DOCKER"
}

locals {
  artifact_repo            = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.docker.name}"
  example_gcp_docker_image = "${local.artifact_repo}/example-gcp"
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
      id   = "docker_build_example"
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t", "${local.example_gcp_docker_image}:$SHORT_SHA",
        "-t", "${local.example_gcp_docker_image}:latest",
        "-f", "./cmd/example-gcp/Dockerfile",
        "."
      ]
      wait_for = ["-"]
    }
    step {
      name = "gcr.io/cloud-builders/docker"
      args = [
        "push",
        "--all-tags",
        local.example_gcp_docker_image
      ]
      wait_for = ["docker_build_example"]
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
