output "log_bucket" {
  description = "Log GCS bucket"
  value       = google_storage_bucket.log_bucket
}
