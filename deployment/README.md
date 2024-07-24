# Deployment

This directory contains configuration-as-code to deploy Trillian Tessera to supported infrastructure:
 - `modules`: terraform modules to configure infrastructure for running a Tessera log.
   + `gcp`: a Tessera GCP specific terraform module.
 - `live`: example terragrunt configurations for deploying to different environments which use the modules

## Prerequisites

Deploying these examples requires installation of:
 - [`terraform`](https://developer.hashicorp.com/terraform/install) or 
   [`opentofu`](https://opentofu.org/docs/intro/install/)
 - [`terragrunt`](https://terragrunt.gruntwork.io/docs/getting-started/install/)

## Deploying

See individual `live` subdirectories.

