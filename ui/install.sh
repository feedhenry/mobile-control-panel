#!/bin/sh

set -x
set -e

SCRIPT_PATH=$(dirname $0)
SCRIPT_ABSOLUTE_PATH=$(cd $SCRIPT_PATH && pwd)

# Mobile API Server install script location
MOBILE_API_SERVER_DIR=$SCRIPT_ABSOLUTE_PATH/../server
MOBILE_API_SERVER_INSTALL_SCRIPT=$MOBILE_API_SERVER_DIR/hack/install-apiserver/openshift/install.sh

# master-config.yaml location
OPENSHIFT_CONFIG_DIR=$SCRIPT_ABSOLUTE_PATH
OPENSHIFT_MASTER_CONFIG=$OPENSHIFT_CONFIG_DIR/master/master-config.yaml
MCP_BASE_DIR=$OPENSHIFT_CONFIG_DIR

# Start OpenShift with the current directory as the config dir
# Using the current directory as the config dir is important for extension development
# This allows changing of extension files and having them already mounted in the origin container
oc cluster up --service-catalog --host-config-dir=${OPENSHIFT_CONFIG_DIR} --use-existing-config=true --version='v3.6.0-rc.0'

# Grant unauthenticated access to the template service broker api
oc login -u system:admin
oc adm policy add-cluster-role-to-group system:openshift:templateservicebroker-client system:unauthenticated system:authenticated

# Enable Extension Development
sudo chmod 666 $OPENSHIFT_MASTER_CONFIG
cd $SCRIPT_ABSOLUTE_PATH
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
