output "conformance_url" {
  description = "The URL of the running conformance server"
  value       = google_cloud_run_v2_service.default.uri
}
