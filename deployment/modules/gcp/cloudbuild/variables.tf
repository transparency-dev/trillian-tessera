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

variable "log_origin" {
  description = "The origin string for the conformance log"
  type        = string
}

variable "kms_key_version_id" {
  description = "The resource ID for the (externally created) KMS key version to use for signing checkpoints"
  type        = string
}


