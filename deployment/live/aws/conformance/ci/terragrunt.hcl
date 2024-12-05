include "root" {
  path   = find_in_parent_folders()
  expose = true
}

inputs = include.root.locals
