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
  }
)
