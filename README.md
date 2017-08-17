# mobile-control-panel (mcp)
The mobile control panel aims to help mobile developers working with and integrating with mobile enabled services on OpenShift. It provides a mobile centric view and set of features designed to enable mobile developers to focus on building great mobile applications with deep server side integrations to powerful services such as server side data sync, push notification, authentication etc without necessarily worrying about how to provision and configure those services to work with their mobile clients.

This repo has 3 main parts:

1) ```./server``` The api server which serves as the server side logic for the mcp UI.

2) ```./ui``` The mcp ui. This is an extension to the OpenShift UI to give a mobile centric view and set of features.

3) ```./docs``` This holds design docs, architectural docs, use cases and development guides etc.

## Prerequisites

* `oc` cli >= 3.6.0-rc0 https://github.com/openshift/origin/releases (Known issue with service-catalog in 3.6.0)
* Node.js >= 4(for install script(s))

## Installation

The `installer/` directory contains a collection of Ansible roles to install
the Mobile Control Panel in OpenShift.
The installer will perform a number of tasks:

1. Install the oc tool on the machine
2. Invoke `oc cluster up` with the Service Catalog enabled
3. Adding templates to the Template Service Broker in OpenShift
4. Installing the MCP UI and server components/resources
5. Creating an Ansible Service Broker in OpenShift

### Prerequisites

* A DockerHub account, credentials are required to set up the Ansible Service
Broker.
* Execute `ansible-galaxy install -r requirements.yml` in the current directory to
install dependencies.
* User with sudo permissions on machine.

### Example

**Running the installer against localhost:**

`ansible-playbook playbook.yml -e "dockerhub_username=myuser" -e "dockerhub_password=mypass" --ask-become-pass`

Once Ansible finished run `oc cluster status` to get the URL of the web console.

To create a mobile app using `oc`:

```
oc create -f ./server/hack/install-apiserver/MobileApp.json
```

### Variables

Variables can be provided as arguments to the `ansible-playbook` command using
`-e variable_name=variable_value` or by populating the
`vars/mobile-control-panel.yml` file.

#### Mandatory
* `dockerhub_username` - DockerHub username
* `dockerhub_password` - DockerHub password

***Note: If `dockerhub_username` and `dockerhub_password` are not specified the
Ansible Service broker will not be created.***

#### Optional
* `templates` - A list of templates to create for the Template Service Broker.
The values can be local paths or URLs. Example: `-e "templates=['/path/to/template.yaml']"`
* `host_config_dir` - Where to create or use the existing host config during `oc cluster up`.
* `cluster_version` - The version of images to use during `oc cluster up`.
* `cluster_public_hostname` - The hostname to use with `--public-ip` and `--routing-suffix` options in `oc cluster up`.

## Cleanup

If you want to just stop the cluster:

```
oc cluster down
```

To stop the cluster and remove any openshift config (Destructive):

```bash
./ui/clean.sh
```

## Troubleshooting

### error: unable to recognize "./server/hack/install-apiserver/MobileApp.json": no matches for mobile.k8s.io/, Kind=MobileApp

If you get this error straight after installing locally, the mobile-apiserver may not be running yet.
You can check that by doing:

```
oc get po -l 'app=apiserver' -n mobile

NAME                         READY     STATUS    RESTARTS   AGE
apiserver-1747434594-pzdvv   2/2       Running   0          13m
```

You can debug the reason why its not running by using `oc get events -n mobile` and looking for any errors or failure events.
If the apiserver is not showing, it may have failed to install. Check the install logs for any errors.
