# Mobile Control Panel UI


## Prerequisites

* `oc` cli >= 3.6.0-rc0 https://github.com/openshift/origin/releases (Known issue with service-catalog in 3.6.0)
* Node.js >= 4(for install script(s))

## Local Installation

Start openshift, configure it for MCP Extension development, and install the mobile apiserver

```
./install.sh
```

The MCP Extension should now be visible in the OpenShift Web Console after the 'origin' container has restarted.
Visit https://127.0.0.1:8443 to see the Console.

To create a mobile app, use `oc` e.g.

```
oc create -f ../server/hack/install-apiserver/MobileApp.json
```
