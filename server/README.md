### Mobile Api Server

The mobile api server is an extension to the Kubernetes api in order to represent mobile objects and apis.
It also has non resource apis.

### Install locally on oc cluster up
To install and setup to try out the mcp, follow the instructions in the main Readme at the root of the repo.


### Developing

You will need go 1.7 + installed. Recommend install go 1.8.x

[download and install golang](https://golang.org/dl/)

Currently you will need Kubernetes on your GOPATH. This is to allow us to build the codegen tools that are needed for maintaining the clients for the api. There are changes coming upstream soon that will mean this wont be needed in the future.

```
go get -u k8s.io/kubernetes

cd $GOPATH/src/k8s.io/kubernetes

git checkout v1.6.8
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
│   ├── install-apiserver
│   │   ├── catalog
│   │   ├── kubernetes
│   │   └── openshift
│   └── template-service-broker
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
