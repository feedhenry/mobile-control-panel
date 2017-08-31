#!/bin/bash

readonly DOCKERHUB_USER="${1}"
readonly DOCKERHUB_PASS="${2}"
readonly DOCKERHUB_ORG="${3}"
readonly LAUNCH_APB_ON_BIND="${4}"
readonly ROUTING_SUFFIX="${5}"

oc login -u system:admin
oc new-project ansible-service-broker
oc process -f /tmp/deploy-ansible-service-broker.template.yaml \
    -n ansible-service-broker \
    -p DOCKERHUB_USER="${DOCKERHUB_USER}" \
    -p DOCKERHUB_PASS="${DOCKERHUB_PASS}" \
    -p DOCKERHUB_ORG="${DOCKERHUB_ORG}" \
    -p LAUNCH_APB_ON_BIND="${LAUNCH_APB_ON_BIND}" \
    -p ROUTING_SUFFIX="${ROUTING_SUFFIX}" | oc create -f -
