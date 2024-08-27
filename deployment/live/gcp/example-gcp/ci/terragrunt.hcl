terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//example-gcp"
}

include "root" {
  path   = find_in_parent_folders()
  expose = true
}

inputs = merge(
  include.root.locals,
  {
    example_gcp_docker_image = "us-central1-docker.pkg.dev/trillian-tessera/docker-prod/example-gcp:latest"
    log_origin               = "example-gcp"
  }
)
