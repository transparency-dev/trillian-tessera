variable "project_id" {
  description = "The project ID to host the builds in"
  type        = string
}

variable "region" {
  description = "The region to host the builds in"
  type        = string
}

variable "env" {
  description = "Unique identifier for the env, e.g. ci or prod"
  type        = string
}

variable "service_account" {
  description = "Service account email to use for cloudbuild"
}

