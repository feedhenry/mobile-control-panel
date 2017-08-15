#!/bin/bash

readonly DOCKERHUB_USER="${0}"
readonly DOCKERHUB_PASS="${1}"
readonly DOCKERHUB_ORG="ansibleplaybookbundle"

curl -s https://raw.githubusercontent.com/openshift/ansible-service-broker/master/templates/deploy-ansible-service-broker.template.yaml > /tmp/deploy-ansible-service-broker.template.yaml

oc login -u system:admin
oc new-project ansible-service-broker
oc process -f /tmp/deploy-ansible-service-broker.template.yaml \
    -n ansible-service-broker \
    -p DOCKERHUB_USER="${DOCKERHUB_USER}" \
    -p DOCKERHUB_PASS="${DOCKERHUB_PASS}" \
    -p DOCKERHUB_ORG="${DOCKERHUB_ORG}" | oc create -f -

if [ "${?}" -ne 0 ]; then
	echo "Error processing template and creating deployment"
	exit
fi

ASB_ROUTE=`oc get routes | grep ansible-service-broker | awk '{print $2}'`

cat <<EOF > /tmp/ansible-service-broker.broker
    apiVersion: servicecatalog.k8s.io/v1alpha1
    kind: Broker
    metadata:
      name: ansible-service-broker
    spec:
      url: https://${ASB_ROUTE}
      authInfo:
        basicAuthSecret:
          namespace: ansible-service-broker
          name: asb-auth-secret
EOF

oc create -f /tmp/ansible-service-broker.broker
