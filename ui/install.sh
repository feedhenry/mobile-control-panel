#!/bin/sh

set -x
set -e

# Mobile API Server repo location
MOBILE_API_SERVER_REPO=github.com/feedhenry/mobile-apiserver
MOBILE_API_SERVER_DIR=$GOPATH/src/$MOBILE_API_SERVER_REPO
MOBILE_API_SERVER_INSTALL_SCRIPT=$MOBILE_API_SERVER_DIR/hack/install-apiserver/openshift/install.sh

# Pull down the mobile-api server repo if not already cloned
[[ -d $MOBILE_API_SERVER_DIR ]] || go get -u $MOBILE_API_SERVER_REPO

# Start OpenShift with the current directory as the config dir
# Using the current directory as the config dir is important for extension development
# This allows changing of extension files and having them already mounted in the origin container
oc cluster up --service-catalog --host-config-dir=${PWD} --version='v3.6.0-rc.0'

# Grant unauthenticated access to the template service broker api
oc login -u system:admin
oc adm policy add-cluster-role-to-group system:openshift:templateservicebroker-client system:unauthenticated system:authenticated

# master-config.yaml location
OPENSHIFT_CONFIG_DIR=$PWD
OPENSHIFT_MASTER_CONFIG=$OPENSHIFT_CONFIG_DIR/master/master-config.yaml
MCP_BASE_DIR=$OPENSHIFT_CONFIG_DIR

# Enable Extension Development
sudo chmod 666 $OPENSHIFT_MASTER_CONFIG
npm i
node update_master_config.js $OPENSHIFT_MASTER_CONFIG
sudo chmod 644 $OPENSHIFT_MASTER_CONFIG

# Allow HostDir Volumes
oc patch scc/restricted -p '{"allowHostDirVolumePlugin":true}'

# Install the Mobile API Server
cd $MOBILE_API_SERVER_DIR
bash -x $MOBILE_API_SERVER_INSTALL_SCRIPT

# TODO: Wait for successful start of the Mobile API Server

# Grant unauthenticated access to the Mobile API Server
oc login -u system:admin
oc adm policy add-cluster-role-to-group mobile-api-caller system:unauthenticated system:anonymous system:authenticated

# Restart openshift to pick up master-config.yaml changes for extensions
docker restart origin
