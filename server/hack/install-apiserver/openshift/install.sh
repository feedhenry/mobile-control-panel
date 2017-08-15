#!/bin/bash
#PUBLIC_IP="192.168.37.1"
#HOSTNAME=${PUBLIC_IP}.nip.io
#ROUTING_SUFFIX="${HOSTNAME}"
#sudo ifconfig lo0 alias ${PUBLIC_IP}
#oc cluster up --image=openshift/origin --version=v3.6.0-rc.0 --service-catalog=true --routing-suffix=${ROUTING_SUFFIX} --public-hostname=${HOSTNAME}  --host-config-dir=/Users/kelly/openshift-config/openshift.local.config
#oc cluster up --image=openshift/origin --version=v3.6.0-rc.0 --service-catalog=true --host-config-dir=/Users/kelly/openshift-config/openshift.local.config

set -x
set -e

targetNamespace=mobile
apiserverConfigDir=/tmp/mobile-apiserver/config
scriptPath=$(dirname $0)
scriptAbsolutePath=$(cd $scriptPath && pwd)
openshiftConfigDir=$(cd $scriptAbsolutePath/../../../../ui && pwd)
masterConfigDir=${openshiftConfigDir}/master
mkdir -p ${apiserverConfigDir} || true
#nodeManifestDir=/Users/kelly/openshift-config/openshift.local.config/node-localhost/static-pods


# 1.  Installer creates namespace of its choice
oc create namespace ${targetNamespace}

# 2.  Installer creates `apiserver` service account. This is the account that will run the pod and is also used to authenticate against master. It is referenced in the deployment.
oc  -n ${targetNamespace} create sa apiserver
until oc -n ${targetNamespace} sa get-token apiserver; do
	echo "waiting for oc -n ${targetNamespace} get secrets/api-serving-cert"
	sleep 1
done




# 3. deploy our api server and register the api service with master
# OK for these to exit with non-zero if resrouces already exist
set +e
oc new-app -f hack/install-apiserver/openshift/deployment.json -n ${targetNamespace}
oc create -f hack/install-apiserver/apiservice.json
# 4. set up some policies to allow our apiserver to delegate auth
oc create policybinding kube-system -n kube-system
oc adm policy add-cluster-role-to-user system:auth-delegator -n ${targetNamespace} -z apiserver
oc adm policy add-role-to-user extension-apiserver-authentication-reader -n kube-system --role-namespace=kube-system system:serviceaccount:${targetNamespace}:apiserver
set -e

until oc -n ${targetNamespace} get secrets/api-serving-cert; do
	echo "waiting for oc -n ${targetNamespace} get secrets/api-serving-cert"
	sleep 1
done


# 5. Extract the cert and key from the automatically created secret based on the service annotation "service.alpha.openshift.io/serving-cert-secret-name": "api-serving-cert" . This provides https in cluster that the master trusts using the keys created for the service account
oc -n ${targetNamespace} extract secret/api-serving-cert --to=${apiserverConfigDir}
mv ${apiserverConfigDir}/tls.crt ${apiserverConfigDir}/serving.crt
mv ${apiserverConfigDir}/tls.key ${apiserverConfigDir}/serving.key


# the front-proxy-ca allows the apiserver to trust the user and group headers sent on the the proxy authentication server / master.
cp ${masterConfigDir}/front-proxy-ca.crt ${apiserverConfigDir}/

#setup the server as a broker

oc create -f hack/install-apiserver/catalog/mobile-broker.json