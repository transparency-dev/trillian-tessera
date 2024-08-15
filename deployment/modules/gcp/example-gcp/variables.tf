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

variable "example_gcp_docker_image" {
  description = "The full image URL (path & tag) for the example-gcp Docker image to deploy"
  type        = string
}
