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
    base_name = get_env("TESSERA_BASE_NAME", "ci-conformance-${substr(uuid(), 0, 4)}")
    enable_antispam = true
  }
)
