# GCP Conformance Configs

## Prerequisites

You'll need to have already configured/created whatever service accounts + IAM permissions
you require, and update the terragrunt files here to match.

## Overview

This config uses the [gcp/conformance](/deployment/gcp/conformance) module to define a
conformance testing environment. At a high level, this environment consists of:
- Spanner DB,
- GCS Bucket,
- CloudRun service running the [GCP-specific conformance binary](/cmd/conformance/gcp).

The config allows identities (e.g. service accounts) to be provided to allow access to
reading from, and writing to, the log.

## Automatic deployment

For the most part, this terragrunt config is automatically used as part of conformance
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
export TESSERA_SIGNER={VALUE} # This should be a note signer string
export TESSERA_VERIFIER={VALUE} # This should be a note verifier string, correspoding to the provided signer.
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

