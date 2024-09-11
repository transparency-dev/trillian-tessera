# GCP Conformance Configs

## Prerequisites

You'll need to have already configured/created a KMS key which can safely be used by the
conformance log.

> [Warning]
> This key should not be used elsewhere or be in any way valuable!

## Automatic deployment

For the most part, this terragrunt config is automatically used as part conformance
testing by the [CloudBuild](/deployment/live/cloudbuild) pipeline, so doesn't generally
need to be manually applied.

## Manual deployment 

First authenticate via `gcloud` as a principle with sufficient ACLs for
the project:
```bash
gcloud auth application-default login
```

Set the required environment variables:
```bash
export GOOGLE_PROJECT={VALUE}
export TESSERA_KMS_KEY_VERSION={VALUE} # This should be the resource name of the key version created above
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

