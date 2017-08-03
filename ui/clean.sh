#!/bin/sh

set -x

# Stop OpenShift
oc cluster down

# Remove openshift config dir things
sudo rm -rf ./master
sudo rm -rf ./node-localhost
