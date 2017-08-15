# Mobile Control Panel (MCP) Ansible installer

A collection of Ansible roles to install the Mobile Control Panel in OpenShift.
This project will perform a number of tasks, namely:

* Install the oc tool on the machine
* Invoke `oc cluster up` with the Service Catalog enabled
* Adding templates to the Template Service Broker in OpenShift
* Installing the MCP UI and server components/resources
* Creating an Ansible Service Broker in OpenShift

## Prerequisites

* A DockerHub account, credentials are required to set up the Ansible Service
Broker.
* Execute `ansible-galaxy install -r requirements.yml` in the current directory to
install dependencies.
* User with sudo permissions on machine.

## Example

**Running the installer against localhost:**
`ansible-playbook playbook.yml -e "dockerhub_username=myuser" -e "dockerhub_password=mypass" --ask-become-pass`

Once Ansible finished run `oc cluster status` to get the URL of the web console.

## Variables

Variables can be provided as arguments to the `ansible-playbook` command using
`-e variable_name=variable_value` or by populating the
`vars/mobile-control-panel.yml` file.

### Mandatory
* `dockerhub_username` - DockerHub username
* `dockerhub_password` - DockerHub password

***Note: If `dockerhub_username` and `dockerhub_password` are not specified the
Ansible Service broker will not be created.***

TODO: finish optional variables
