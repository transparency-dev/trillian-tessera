output "log_bucket_id" {
  description = "Log S3 bucket name"
  value       = module.storage.log_bucket.id
}

output "log_bucket_http" {
  description = "Log S3 bucket http access"
  value       = "https://${module.storage.log_bucket.bucket_regional_domain_name}"
}

output "log_rds_db" {
  description = "Log RDS database endpoint"
  value       = module.storage.log_rds_db.endpoint
}

output "log_name" {
  description = "Log name"
  value       = local.name
}
