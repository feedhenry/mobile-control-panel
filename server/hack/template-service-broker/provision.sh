#!/bin/bash

# Provision OpenShift Template Service Broker (potsb)

# In order to enable access to the TSB we must grant unauthenticated access to
# the template service broker api.
readonly GROUP="system:openshift:templateservicebroker-client"
readonly ROLES="system:unauthenticated system:authenticated"
# The 'system:admin' user has permissions to grant unauthenticated access etc.
# so we'll use that user.
readonly USER="system:admin"
# By default, any templates in the 'openshift' project will be exposed in the
# catalog.
readonly PROJECT="openshift"

# We'll be messing with OpenShift users, so let's *try* to return you back to
# your previous state after we complete.
originalUser="$(oc whoami)"
originalProject="$(oc project -q)"

# Change user.
oc login -u "${USER}"

# Change project.
oc project "${PROJECT}"

# Add those roles that are required, mentioned above.
oc adm policy add-cluster-role-to-group "${GROUP}" ${ROLES}

# Go through each argument, they can be local paths or URLs. Some may fail
# because they already exist. That's fine, we'll carry on.
for templateUrl in "${@}"
do
  oc create -f "${templateUrl}"
done

echo -e "\nProvisioning complete. Trying to restore user state...\n"

# Try to restore the original user state.
oc login -u "${originalUser}"
oc project "${originalProject}"

echo -e "\nUser state restored. user=$(oc whoami) project=$(oc project -q)"
echo -e "\nHooray! We're all done! Go to your OpenShift console. You should be
able to see a bunch of templates. If not, refresh, it can take a while."
