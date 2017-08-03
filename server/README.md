### Mobile Api Server

The mobile api server is an extension to the Kubernetes api in order to represent mobile objects and apis.

### Install locally on oc cluster up
After cloning this repo take the following steps:

*Note* improvements will be made to this over time. It is currently just a POC

- Using v3.6.0-rc0 https://github.com/openshift/origin/releases/tag/v3.6.0-rc.0 . Start your oc cluster with a host config and service catalog enabled. We enabled the service catalog so that Kubernetes and OpenShift are set up to recognise the ```apiregistration.k8s.io``` api.

```
oc cluster up --service-catalog --host-config-dir=${HOME}/openshift-config/openshift.local.config
```
- Login as system:admin

```
oc login -u system:admin
```

- Change the restricted scc to allow hostPath volumes. This is so we can mount the front-proxy-ca.crt which is needed for authentication. Plan to change how this is done in the future.

```
oc edit scc restricted

change: 

allowHostDirVolumePlugin:false 

to 

allowHostDirVolumePlugin:true

```

- Run the install script. Note there will likely be an error about roles already existing. This is fine.

```
./hack/install-apiserver/openshift/install.sh
```

- Check the api is set up. You should be able to create MobileApps and get MobileApps

```
oc login 
developer
developer

oc create -f ./hack/install-apiserver/MobileApp.json

oc get mobileapps
```


### Developing

You will need go 1.7 + installed. Recommend install go 1.8.x

[download and install golang](https://golang.org/dl/)

Currently you will need Kubernetes on your GOPATH. This is to allow us to build the codegen tools that are needed for maintaining the clients for the api. There are changes coming upstream soon that will mean this wont be needed in the future.

```
go get -u k8s.io/kubernetes

cd $GOPATH/src/k8s.io/kubernetes

git checkout v1.7.0
```

Dependencies are managed by dep

```
go get -u github.com/golang/dep/cmd/dep
```

Clone the root repo into your $GOPATH/src/github.com/feedhenry/mobile-control-panel

```
mkdir -p $GOPATH/src/github.com/feedhenry

cd $GOPATH/src/github.com/feedhenry

git clone git@github.com:feedhenry/mobile-control-panel.git

```



you can then you the Makefile to do your first build:

```
$GOPATH/src/github.com/feedhenry/mobile-control-panel/server

make images
```

### Directory Structure

```
├── artifacts
│   └── image
├── cmd
│   ├── api_test #simple cli for hitting api during local dev
│   └── server #main artifact
│       ├── apiserver.local.config
│       │   └── certificates
│       └── start
├── hack # various scripts for use during development
│   └── install-apiserver
│       ├── kubernetes
│       └── openshift
├── pkg # main go packages
│   ├── apis # the api resource definitions and registration such as MobileApp
│   │   └── mobile
│   │       ├── install
│   │       └── v1alpha1
│   ├── apiserver # api server setup
│   ├── client #generated clients for talking to etcd and the master api server
│   │   ├── clientset_generated
│   │   │   ├── clientset
│   │   │   │   ├── fake
│   │   │   │   ├── scheme
│   │   │   │   └── typed
│   │   │   │       └── mobile
│   │   │   │           └── v1alpha1
│   │   │   │               └── fake
│   │   │   └── internalclientset
│   │   │       ├── fake
│   │   │       ├── scheme
│   │   │       └── typed
│   │   │           └── mobile
│   │   │               └── internalversion
│   │   │                   └── fake
│   │   ├── informers_generated
│   │   │   ├── externalversions
│   │   │   │   ├── internalinterfaces
│   │   │   │   └── mobile
│   │   │   │       └── v1alpha1
│   │   │   └── internalversion
│   │   │       ├── internalinterfaces
│   │   │       └── mobile
│   │   │           └── internalversion
│   │   └── listers_generated
│   │       └── mobile
│   │           ├── internalversion
│   │           └── v1alpha1
│   ├── mobile #additional logic for mobile
│   └── registry # etcd logic and validation
│       └── mobileapp
└── test
    └── integration
        └── apiserver.local.config
            └── certificates
```            
