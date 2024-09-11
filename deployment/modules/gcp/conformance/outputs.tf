output "run_service_account" {
  description = "The CloudRun service account"
  value       = google_service_account.cloudrun_service_account
}

output "conformance_url" {
  description = "The URL of the running conformance server"
  value       = google_cloud_run_v2_service.default.uri
}
