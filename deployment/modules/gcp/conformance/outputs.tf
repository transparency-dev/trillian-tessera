output "conformance_url" {
  description = "The URL of the running conformance server"
  value       = google_cloud_run_v2_service.default.uri
}

output "conformance_bucket_name" {
  description = "The name of the conformance log bucket"
  value       = module.gcs.log_bucket.name
}
