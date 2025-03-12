variable "project_id" {
  description = "GCP project ID where the log is hosted"
  type        = string
}

variable "base_name" {
  description = "Base name to use when naming resources"
  type        = string
}

variable "location" {
  description = "Location in which to create resources"
  type        = string
}

variable "env" {
  description = "Environment name, e.g ci, prod, etc."
  type        = string
}

variable "server_docker_image" {
  description = "The full image URL (path & tag) for the Docker image to deploy in Cloud Run"
  type        = string
}

variable "signer" {
  description = "The note signer which should be used to sign checkpoints"
  type        = string
}

variable "cloudrun_service_account" {
  description = "The service account email to use for the CloudRun instance. If unset, uses the project default service account."
  type        = string
}

variable "conformance_writers" {
  description = "The list of users allowed to invoke HTTP calls to the conformance Cloud Run instance. If unset, only the project default service account will be able to send requests."
  type        = list(any)
}

variable "conformance_readers" {
  description = "The list of users allowed to read the conformance t-log resources from GCS. If unset, only the project default service account will be able to read the t-log contents."
  type        = list(any)
}

variable "enable_antispam" {
  description = "Set to true to enable persistent antispam feature on conformance instance."
  type        = bool
}
