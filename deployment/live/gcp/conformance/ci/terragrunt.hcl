terraform {
  source = "${get_repo_root()}/deployment/modules/gcp//conformance"
}

include "root" {
  path   = find_in_parent_folders()
  expose = true
}

inputs = merge(
  include.root.locals,
  {
    server_docker_image = "us-central1-docker.pkg.dev/trillian-tessera/docker-prod/conformance-gcp:latest"
  }
)
