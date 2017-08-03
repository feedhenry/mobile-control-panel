# API Server


## Why a Kubernetes based API Server
The API Server is based on a standard (Kubernetes API Server)[https://kubernetes.io/docs/admin/kube-apiserver/] model. As the mobile control panel is *heavily* integrated into OpenShift/Kubernetes, we want to leverage as much of existing infrastructure as possible. 

In Kubernetes 1.6 and 1.7 the ability to extend the Kubernetes API was introduced and with it came examples of how to setup an (Extension API Server)[https://kubernetes.io/docs/tasks/access-kubernetes-api/setup-extension-api-server/]. This means that by using the standard Kubernetes apiserver packages and setup as is used by the core Kubernetes project, not only can we extend the API with our own objects but it also allows us to leverage and build upon a battle tested and hardened, standard feature set:
- Authentication and authorization is delegated to Kubernetes core, we have no need to re write how authentication and authorisation will work or a custom integration layer.
- Audit logging is built in as part of the standard api server
- The default REST implementations for crudl of the resource are ready made and present from the start with hooks provided for adding custom validation.
- Multi-tenanted out of the box when running in OpenShift cluster
- API versioning is built in.
- API discovery is built in. This means any resource we add can automatically be used via the oc tool.
- Secure serving as standard
- Large open source community committed to making Kubernetes extendable
- Well maintained client library that can be used to leverage the rest the (Kubernetes API)[https://github.com/kubernetes/client-go]

## Responsibilities
The mobile API server provides the REST API used by the mobile control panel. 

- Storing, validating and managing the MobileApp resource. 
- Acting on various state changes in the namespaces where the MobileApp was created. (for example when bindings to services are created as secrets it will use them to construct the mobile sdk configuration).
- Adding custom APIs to perform various business logic based on requests made by the MCP (example generating and downloading the sdk configuration file)

## Backing Data Store
The API server has a lot of existing infrastructure to make use of the etcd key value distributed store. So making used of etcd makes a lot of sense. Also the hope is that we can connect to the master etcd provisioned with OpenShift and use that as the backing store.
