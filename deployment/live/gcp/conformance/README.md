## Deployment 

First authenticate via `gcloud` as a principle with sufficient ACLs for
the project:
```bash
gcloud auth application-default login
```

Set your GCP project ID with:
```bash
export GOOGLE_PROJECT={VALUE}
```

Optionally, customize the GCP region (defaults to "us-central1"),
and bucket name prefix (defaults to "conformance"):
```bash
export GOOGLE_REGION={VALUE}
export TESSERA_BASE_NAME={VALUE}
```

Terraforming the project can be done by:
 1. `cd` to the relevant directory for the environment to deploy/change (e.g. `ci`)
 2. Run `terragrunt apply`

