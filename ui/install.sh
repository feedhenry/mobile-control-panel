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
$MOBILE_API_SERVER_INSTALL_SCRIPT
