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
  description   = "Tessera conformance docker images"
  format        = "DOCKER"
}

locals {
  artifact_repo                = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.docker.name}"
  conformance_gcp_docker_image = "${local.artifact_repo}/conformance-gcp"
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
    /*
    step {
      id = "docker_build_conformance_gcp"
      name = "gcr.io/cloud-builders/docker"
      args = [
        "build",
        "-t", "${local.conformance_gcp_docker_image}:$SHORT_SHA",
        "-t", "${local.conformance_gcp_docker_image}:latest",
        "-f", "./cmd/conformance/gcp/Dockerfile",
        "."
      ]
    }
    step {
      id   = "docker_push_conformance_gcp"
      name = "gcr.io/cloud-builders/docker"
      args = [
        "push",
        "--all-tags",
        local.conformance_gcp_docker_image
      ]
      wait_for = ["docker_build_conformance_gcp"]
    }
    */
    step {
      id         = "terraform_apply_conformance_ci"
      name       = "alpine/terragrunt"
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
    //  wait_for = ["docker_push_conformance_gcp"]
    }
    step {
      id = "terraform_outputs"
      name = "alpine/terragrunt"
      script = <<EOT
        cd deployment/live/gcp/conformance/ci
        terragrunt output --raw conformance_url > /workspace/conformance_url
      EOT
      wait_for = ["terraform_apply_conformance_ci"]
    }
    /*
    step {
      id   = "generate_verifier"
      name = "golang"
      args = [
        "go",
        "run",
        "./cmd/conformance/gcp/kmsnote",
        "--key_id=${var.kms_key_version_id}",
        "--name=${var.log_origin}",
        "--output=/workspace/verifier.pub"
      ]
    }
    */
    step { 
      id = "access"
      name = "gcr.io/cloud-builders/gcloud"
      script = <<EOT
      gcloud auth print-access-token > /workspace/cb_access
      curl -H "Metadata-Flavor: Google" "http://metadata/computeMetadata/v1/instance/service-accounts/${google_service_account.cloudbuild_service_account.email}/identity?audience=$(cat /workspace/conformance_url)" > /workspace/cb_identity
      cat /workspace/cb_identity
      EOT
      wait_for = ["terraform_outputs"]
    }
    step {
      id   = "hammer"
      name = "golang"
      script = <<EOT
      go run ./hammer --log_public_key=ci-conformance+1e5feae8+ARe2PncWChrXMrQtDOqJrKkn2XXOGwsJ+amwLnaWyhEK --log_url=https://storage.googleapis.com/trillian-tessera-ci-conformance-bucket/ --write_log_url="$(cat /workspace/conformance_url)" -v=2 --show_ui=false --bearer_token="$(cat /workspace/cb_access)" --bearer_token_write="$(cat /workspace/cb_identity)" --logtostderr --num_writers=1500 --max_write_ops=2000 --leaf_min_size=1024 --leaf_write_goal=100000
      EOT
      wait_for = ["terraform_outputs", "access"]
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

