# Services
resource "google_project_service" "serviceusage_googleapis_com" {
  service            = "serviceusage.googleapis.com"
  disable_on_destroy = false
}
resource "google_project_service" "storage_api_googleapis_com" {
  service            = "storage-api.googleapis.com"
  disable_on_destroy = false
}
resource "google_project_service" "storage_component_googleapis_com" {
  service            = "storage-component.googleapis.com"
  disable_on_destroy = false
}
resource "google_project_service" "storage_googleapis_com" {
  service            = "storage.googleapis.com"
  disable_on_destroy = false
}

## Resources

# Buckets

resource "google_storage_bucket" "log_bucket" {
  name                        = "${var.project_id}-${var.base_name}-bucket"
  location                    = var.location
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  force_destroy               = var.ephemeral
}

resource "google_storage_bucket_iam_binding" "log_bucket_reader" {
  bucket = google_storage_bucket.log_bucket.name
  role   = "roles/storage.objectViewer"
  members = var.bucket_readers
}

resource "google_storage_bucket_iam_binding" "log_bucket_writer" {
  bucket  = google_storage_bucket.log_bucket.name
  role    = "roles/storage.legacyBucketWriter"
  members = var.log_writer_members
}

# Spanner

resource "google_spanner_instance" "log_spanner" {
  name             = var.base_name
  config           = "regional-${var.location}"
  display_name     = var.base_name
  processing_units = 100

  force_destroy    = var.ephemeral
}

resource "google_spanner_database" "log_db" {
  instance = google_spanner_instance.log_spanner.name
  name     = "${var.base_name}-db"
  ddl = [
    "CREATE TABLE SeqCoord (id INT64 NOT NULL, next INT64 NOT NULL,) PRIMARY KEY (id)",
    "CREATE TABLE Seq (id INT64 NOT NULL, seq INT64 NOT NULL, v BYTES(MAX),) PRIMARY KEY (id, seq)",
    "CREATE TABLE IntCoord (id INT64 NOT NULL, seq INT64 NOT NULL,) PRIMARY KEY (id)",
  ]

  deletion_protection         = !var.ephemeral
}

resource "google_spanner_database_iam_binding" "database" {
  instance = google_spanner_instance.log_spanner.name
  database = google_spanner_database.log_db.name
  role     = "roles/spanner.databaseUser"

  members = var.log_writer_members
}
