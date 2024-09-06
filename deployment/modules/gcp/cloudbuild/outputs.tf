output "artifact_registry_id" {
  description = "The ID of the created artifact registry for docker images"
  value       = google_artifact_registry_repository.docker.id
}

output "artifact_registry_name" {
  description = "The name of the created artifact registry for docker images"
  value       = google_artifact_registry_repository.docker.name
}

output "cloudbuild_trigger_id" {
  description = "The ID of the created trigger for building images"
  value       = google_cloudbuild_trigger.docker.id
}

output "conformance_gcp_docker_image" {
  description = "The address of the GCP conformance docker image that will be built"
  value       = local.conformance_gcp_docker_image
}
