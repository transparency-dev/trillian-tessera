include "root" {
  path   = find_in_parent_folders()
  expose = true
}

inputs = merge(
  include.root.locals,
  {
    # This hack makes it so that the antispam tables are created in the main
    # tessera DB. We strongly recommend that the antispam DB is separate, but
    # creating a second DB from Terraform is too difficult without a large
    # rewrite. For CI purposes, testing antispam, even if in the same DB, is
    # preferred compared to not testing antispam at all.
    antispam         = true
    antispam_db_name = "tessera"
    create_antispam  = false
  }
)
