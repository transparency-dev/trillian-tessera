## Deployment 

First authenticate via `gcloud` as a principle with sufficient ACLs for
the project:
```bash
gcloud auth application-default login
```

Then, specify your Google Cloud project ID:
```bash
export GOOGLE_PROJECT={VALUE}
```

Eventually, customize the region (defaults to "us-central1"), and bucket name prefix
(defaults to "tessera-example"):
```bash
export GOOGLE_REGION={VALUE}
export TESSERA_BASE_NAME={VALUE}
```

Terraforming the project can be done by:
 1. `cd` to the relevant `live` directory for the environment to deploy/change
 2. Run `terragrunt apply`

