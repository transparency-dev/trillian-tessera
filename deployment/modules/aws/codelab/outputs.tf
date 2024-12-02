output "log_bucket_id" {
  description = "Log S3 bucket name"
  value       = module.storage.log_bucket.id
}

output "log_rds_db" {
  description = "Log RDS database endpoint"
  value       = module.storage.log_rds_db.endpoint
}
