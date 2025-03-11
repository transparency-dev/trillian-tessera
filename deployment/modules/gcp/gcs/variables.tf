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
  description = "Unique identifier for the env, e.g. ci or prod"
  type        = string
}

variable "bucket_readers" {
  description = "List of identities allowed to read the log bucket"
  type        = list(any)
  default     = ["allUsers"]
}

variable "log_writer_members" {
  description = "List of identities in member format allowed to write to the log"
  type        = list(any)
}

variable "create_antispam" {
  description = "Set to true to create the infrastructure required by the GCP antispam implementation"
  type        = bool
}

variable "ephemeral" {
  description = "Set to true if this is a throwaway/temporary log instance. Will set attributes on created resources to allow them to be disabled/deleted more easily."
  type        = bool
}
