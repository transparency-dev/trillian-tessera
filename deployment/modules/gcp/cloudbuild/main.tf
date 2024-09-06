terraform {
  backend "gcs" {}
}

provider "google" {
  project = var.project_id
  region  = var.region
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
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t", "${local.example_gcp_docker_image}:$SHORT_SHA",
        "-t", "${local.example_gcp_docker_image}:latest",
        "-f", "./cmd/example-gcp/Dockerfile",
        "."
      ]
    }
    step {
      name = "gcr.io/cloud-builders/docker"
      args = [
        "push",
        "--all-tags",
        local.example_gcp_docker_image
      ]
      wait_for = ["docker_build_conformance_gcp"]
    }
    step {
      id   = "terraform_apply_conformance_ci"
      name = "alpine/terragrunt"
      entrypoint = "terragrunt"
      args = [
        "--terragrunt-non-interactive",
        "apply",
        "-auto-approve",
      ]
      dir = "deployment/live/gcp/conformance/ci"
      env = [
        "GOOGLE_PROJECT=${var.project_id}",
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

# roles managed externally.
resource "google_service_account" "cloudbuild_service_account" {
  account_id   = "cloudbuild-${var.env}-sa"
  display_name = "Service Account for CloudBuild (${var.env})"
}

