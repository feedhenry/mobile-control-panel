# Provision OpenShift Template Service Broker
A script to provision the Template Service Broker in OpenShift and add any
templates you might want there. Template URLs or local paths can be specified
as a list of arguments to the script.

## Prerequisites

Have an OpenShift cluster >= 3.6 with the Service Catalog enabled.

The most simple way to do this is by running `oc cluster up --service-catalog`.

Once the cluster is up you are ready to run `provision.sh`.

## What does the script do?

1. Logs in to OpenShift as `system:admin`.
2. Uses the `openshift` project.
3. Allow unauthenticated access to the Template Service Broker API.
4. Create each template specified in the arguments in the `openshift` project.

## Examples

### Enable unauthenticated access to the Template Service Broker API
***Note: This will not add any custom templates***
`./provision.sh`

**Result:** You will now be able to see all default availble templates in the
Catalog.

### Add the fh-sync-service template
`./provision.sh https://raw.githubusercontent.com/feedhenry/fh-sync-server/master/fh-sync-server-DEVELOPMENT.yaml`

**Result:** You will now be able to see all default templates along with a Sync
template in the Catalog.
