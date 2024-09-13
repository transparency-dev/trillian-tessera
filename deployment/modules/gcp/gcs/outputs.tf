output "log_bucket" {
  description = "Log GCS bucket"
  value       = google_storage_bucket.log_bucket
}

output "log_spanner_db" {
  description = "Log Spanner database"
  value       = google_spanner_database.log_db
}

output "log_spanner_instance" {
  description = "Log Spanner instance"
  value       = google_spanner_instance.log_spanner
}

