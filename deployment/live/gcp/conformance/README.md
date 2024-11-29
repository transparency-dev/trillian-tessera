# GCP Conformance Configs

This config uses the [gcp/conformance](/deployment/modules/gcp/conformance) module to
define a conformance testing environment. At a high level, this environment consists of:
- Spanner DB,
- GCS Bucket,
- CloudRun service running the [GCP-specific conformance binary](/cmd/conformance/gcp).

The config allows identities (e.g. service accounts) to be provided to allow access to
reading from, and writing to, the log.

## Automatic deployment

For the most part, this terragrunt config is automatically used as part of conformance
testing by the [CloudBuild](/deployment/live/gcp/cloudbuild) pipeline, so doesn't generally
need to be manually applied.

## Manual deployment 

### Prerequisites

You'll need the following tools installed:

- [`docker`](https://docs.docker.com/engine/install/)
- [`gcloud`](https://cloud.google.com/sdk/docs/install)
- One of:
   + [`terraform`](https://developer.hashicorp.com/terraform/install) or
   + [`opentofu`](https://opentofu.org/docs/intro/install/)
- [`terragrunt`](https://terragrunt.gruntwork.io/docs/getting-started/install/)


### Process

First ensure you've got the right Google Cloud project configured, and authenticate via `gcloud`
as a principle with sufficient ACLs for the project:
```bash
gcloud config get project
gcloud config set project {YOUR PROJECT}
gcloud auth application-default login
```

You will need to build and push the image created by [/cmd/conformance/gcp/Dockerfile] somewhere.
Google Artifact Registry is one option: https://cloud.google.com/build/docs/build-push-docker-image

Note that it's not currently possible to _build_ the docker image with Google Cloud Build (this is because
building from the root directory but referencing Dockerfile in a subdirectory isn't supported), but you can
configure your local `docker` to get access to Artifact Registry using the gcloud CLI credential helper, and
then build locally and push to Artifact Registry. Details on this can be found in the link above.

```bash
$ gcloud artifacts repositories create ${DOCKER_REPO_NAME} \
        --repository-format=docker \
        --location=us-central1 \
        --description="My Tessera docker repo" \
        --immutable-tags

$ gcloud auth configure-docker us-central1-docker.pkg.dev

$ docker build . -f ./cmd/conformance/gcp/Dockerfile --tag us-central1-docker.pkg.dev/${YOUR_GCP_PROJECT}/${DOCKER_REPO_NAME}/conformance:latest

$ docker push us-central1-docker.pkg.dev/${YOUR_GCP_PROJECT}/${DOCKER_REPO_NAME}/conformance:latest

```

Set the required environment variables:
```bash
# The ID of the Google Cloud Project you're deploying into
export GOOGLE_PROJECT={VALUE}

# This should be a note signer string
export TESSERA_SIGNER={VALUE}

# A docker image built from /cmd/conformance/gcp/Dockerfile
# Use the name/tag from the docker image you built above.
export TESSERA_CLOUD_RUN_DOCKER_IMAGE=us-central1-docker.pkg.dev/${YOUR_GCP_PROJECT}/${DOCKER_REPO_NAME}/conformance:latest

```

Optionally, set the envrionment variables below to customize the deployment:
```bash
# GCP region to deply into (defaults to us-central1)
export GOOGLE_REGION={VALUE} 

# This is used as part of resource names, using this variable will allow you to have multiple deployments in a single project.
export TESSERA_BASE_NAME={VALUE} 

# This allows you to specify the email of an existing service account which should be used by Cloud Run.
# By default, the project's default service account will be used.
export TESSERA_CLOUD_RUN_SERVICE_ACCOUNT={VALUE}

# This allows configuration of which users are allowed to read from the GCS bucket containing the t-log tiles.
# To make the bucket public, set this to "allUsers".
export TESSERA_READER={VALUE}

# This allows configuration of which users are allowed to make HTTP requests to the Cloud Run instance, e.g. to add entries to the t-log.
# By default, only the project's default service account is permitted.
export TESSERA_WRITER={VALUE}
```

Finally, apply the config using `terragrunt`:
 1. `cd` to the relevant directory for the environment to deploy/change (e.g. `ci`)
 2. Run `terragrunt apply`

