# devconsole

This repository was initially bootstrapped using [CoreOS operator](https://github.com/operator-framework/operator-sdk). 

## Build

### Pre-requisites
- [operator-sdk v0.5.0](https://github.com/operator-framework/operator-sdk#quick-start) 
- [dep][dep_tool] version v0.5.0+.
- [git][git_tool]
- [go][go_tool] version v1.10+.
- [docker][docker_tool] version 17.03+.
- [kubectl][kubectl_tool] version v1.11.0+ or [oc] version 3.11
- Access to a kubernetes v.1.11.0+ cluster or openshift cluster version 3.11

### Build
```
make build
```
## Deployment

### Set up Minishift (one-off)
* create a new profile to test the operator
```
minishift profile set devconsole
```
* enable the admin-user add-on
```
minishift addon enable admin-user
```
* optionally, configure the VM 

```
minishift config set cpus 4
minishift config set memory 8GB
minishift config set vm-driver virtualbox
```
* start the instance
```
minishift start
```
> NOTE: this setup should be deprecated in favor of [OCP4 install]().

### Deploy the operator

In dev mode, simply run your operator locally:
```
make local
```
> NOTE: To watch all namespaces, `APP_NAMESPACE` is set to empty string. 
If a specific namespace is provided only that project will watched. 
As we reuse `openshift`'s imagestreams for build, we need to access all namespaces.

### Deploy the CR for testing
* Make sure minishift is running and use myproject
```
oc project myproject
```
* Clean previously created resources
```
make deploy-clean
```
* Deploy CR
```
make deploy-test
```
* See the newly created resources
```
oc get is,bc,svc,component.devconsole,build
NAME                                           DOCKER REPO                               TAGS      UPDATED
imagestream.image.openshift.io/myapp-output    172.30.1.1:5000/myproject/myapp-output
imagestream.image.openshift.io/myapp-runtime   172.30.1.1:5000/myproject/myapp-runtime   latest    46 seconds ago

NAME                                      TYPE      FROM         LATEST
buildconfig.build.openshift.io/myapp-bc   Source    Git@master   1

NAME                                         AGE
component.devconsole.openshift.io/myapp   48s

NAME                                  TYPE      FROM          STATUS    STARTED          DURATION
build.build.openshift.io/myapp-bc-1   Source    Git@85ac14e   Running   45 seconds ago
```

## Directory layout

Please consult [the documentation](https://github.com/operator-framework/operator-sdk/blob/master/doc/project_layout.md) in order to learn about this project's structure: 

|File/Folders  |Purpose |
|--------------|--------|
| cmd          | Contains `manager/main.go` which is the main program of the operator. This instantiates a new manager which registers all custom resource definitions under `pkg/apis/...` and starts all controllers under `pkg/controllers/...`.|
| pkg/apis | Contains the directory tree that defines the APIs of the Custom Resource Definitions(CRD). Users are expected to edit the `pkg/apis/<group>/<version>/<kind>_types.go` files to define the API for each resource type and import these packages in their controllers to watch for these resource types.|
| pkg/controller | This pkg contains the controller implementations. Users are expected to edit the `pkg/controller/<kind>/<kind>_controller.go` to define the controller's reconcile logic for handling a resource type of the specified `kind`.|
| build | Contains the `Dockerfile` and build scripts used to build the operator.|
| deploy | Contains various YAML manifests for registering CRDs, setting up [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/), and deploying the operator as a Deployment.|
| Gopkg.toml Gopkg.lock | The [dep](https://github.com/golang/dep) manifests that describe the external dependencies of this operator.|
| vendor | The golang [Vendor](https://golang.org/cmd/go/#hdr-Vendor_Directories) folder that contains the local copies of the external dependencies that satisfy the imports of this project. [dep](https://github.com/golang/dep) manages the vendor directly.|


## Enabling the DevOps perspective in OpenShift

The frontend can check for the presence of the devconsole CRDs using the Kubernetes API.  Check for [the existence of a Custom Resource Definitions](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#list-customresourcedefinition-v1beta1-apiextensions) with name as `gitsources.devconsole.openshift.io`.  If it exists, it will enable the DevOps perspective in the Openshift Console.

To install the DevOps Console operator and run it using OLM

```
kubectl create -f https://raw.githubusercontent.com/operator-framework/operator-lifecycle-manager/master/deploy/upstream/quickstart/olm.yaml

kubectl create -f http://operator-hub-shbose-preview1-stage.b542.starter-us-east-2a.openshiftapps.com/install/devconsole.v0.1.0.yaml
```

A `CatalogSource` followed by a new `Subscription` is created when the above commands are executed.

```
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: rhd-operatorhub-catalog
  namespace: olm
spec:
  sourceType: grpc
  image: sbose78/operator-registry:latest
  displayName: Community Operators
  publisher: RHD Operator Hub
---
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: my-devconsole
  namespace: operators
spec:
  channel: alpha
  name: devconsole
  source: rhd-operatorhub-catalog
  sourceNamespace: olm
```

[dep_tool]:https://golang.github.io/dep/docs/installation.html
[git_tool]:https://git-scm.com/downloads
[go_tool]:https://golang.org/dl/
[docker_tool]:https://docs.docker.com/install/
[kubectl_tool]:https://kubernetes.io/docs/tasks/tools/install-kubectl/

