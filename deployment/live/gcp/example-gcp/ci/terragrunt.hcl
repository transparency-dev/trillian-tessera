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
    example_gcp_docker_image = "todo"
  }
)
