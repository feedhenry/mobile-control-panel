# mobile-control-panel (mcp)
The mobile control panel aims to help mobile developers working with and integrating with mobile enabled services on OpenShift. It provides a mobile centric view and set of features designed to enable mobile developers to focus on building great mobile applications with deep server side integrations to powerful services such as server side data sync, push notification, authentication etc without necessarily worrying about how to provision and configure those services to work with their mobile clients.

This repo has 3 main parts:

1) ```./server``` The api server which serves as the server side logic for the mcp UI. 

2) ```./ui``` The mcp ui. This is an extension to the OpenShift UI to give a mobile centric view and set of features.

3) ```./docs``` This holds design docs, architectural docs, use cases and development guides etc.

## Prerequisites

* `oc` cli >= 3.6.0-rc0 https://github.com/openshift/origin/releases (Known issue with service-catalog in 3.6.0)
* Node.js >= 4(for install script(s))

## Local Installation

Start openshift, configure it for MCP Extension development, and install the mobile apiserver

```
./ui/install.sh
```

The MCP Extension should now be visible in the OpenShift Web Console after the 'origin' container has restarted.
Visit https://127.0.0.1:8443 to see the Console.

To create a mobile app using `oc`:

```
oc create -f ./server/hack/install-apiserver/MobileApp.json
```

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
If the apiserver is not showing, it may have failed to install. Check the install logs for any errors.Test change
