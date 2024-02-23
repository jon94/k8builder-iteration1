# README
- The controller checks for demo volume crd for spec.Name.
- It compares with the demo volume crd's status.Name, if they are not the same, then it copies spec.Name to status.Name 
- This is a basic exploration of kubebuilder.

# DEMO
```
jonathan.lim@COMP-PW290173CK k8builder-iteration1 % k apply -f config/crd/bases/demo.demo.jonlimpw.io_demovolumes.yaml 
customresourcedefinition.apiextensions.k8s.io/demovolumes.demo.demo.jonlimpw.io created
```
```jonathan.lim@COMP-PW290173CK k8builder-iteration1 % k get crd|grep -i jonlimpw
demovolumes.demo.demo.jonlimpw.io                                  2024-02-23T13:33:15Z
```
```
jonathan.lim@COMP-PW290173CK k8builder-iteration1 % k apply -f config/samples/demo_v1_demovolume.yaml 
demovolume.demo.demo.jonlimpw.io/demovolume-sample created
```
```
jonathan.lim@COMP-PW290173CK k8builder-iteration1 % k get demovolume
NAME                AGE
demovolume-sample   17s
```
```
jonathan.lim@COMP-PW290173CK k8builder-iteration1 % k get demovolume -o yaml
apiVersion: v1
items:
- apiVersion: demo.demo.jonlimpw.io/v1
  kind: DemoVolume
  metadata:
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"demo.demo.jonlimpw.io/v1","kind":"DemoVolume","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"demo","app.kubernetes.io/instance":"demovolume-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"demovolume","app.kubernetes.io/part-of":"demo"},"name":"demovolume-sample","namespace":"default"},"spec":{"Name":"demo","Size":10}}
    creationTimestamp: "2024-02-23T13:33:57Z"
    generation: 1
    labels:
      app.kubernetes.io/created-by: demo
      app.kubernetes.io/instance: demovolume-sample
      app.kubernetes.io/managed-by: kustomize
      app.kubernetes.io/name: demovolume
      app.kubernetes.io/part-of: demo
    name: demovolume-sample
    namespace: default
    resourceVersion: "146251984"
    uid: eec3987c-5599-48ac-bcd7-dafaec462fbe
  spec:
    Name: demo
    Size: 10
kind: List
metadata:
  resourceVersion: ""
```
```
jonathan.lim@COMP-PW290173CK k8builder-iteration1 % make run
/Users/jonathan.lim/k8builder-iteration1/bin/controller-gen-v0.14.0 rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/Users/jonathan.lim/k8builder-iteration1/bin/controller-gen-v0.14.0 object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go run ./cmd/main.go
2024-02-23T21:34:42+08:00       INFO    setup   starting manager
2024-02-23T21:34:42+08:00       INFO    starting server {"kind": "health probe", "addr": "[::]:8081"}
2024-02-23T21:34:42+08:00       INFO    controller-runtime.metrics      Starting metrics server
2024-02-23T21:34:42+08:00       INFO    controller-runtime.metrics      Serving metrics server  {"bindAddress": ":8080", "secure": false}
2024-02-23T21:34:42+08:00       INFO    Starting EventSource    {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume", "source": "kind source: *v1.DemoVolume"}
2024-02-23T21:34:42+08:00       INFO    Starting Controller     {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume"}
2024-02-23T21:34:43+08:00       INFO    Starting workers        {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume", "worker count": 1}
2024-02-23T21:34:43+08:00       INFO    Enter Reconcile {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume", "DemoVolume": {"name":"demovolume-sample","namespace":"default"}, "namespace": "default", "name": "demovolume-sample", "reconcileID": "c92daf9c-4776-4c60-bf1a-e49f2879555a", "req": {"name":"demovolume-sample","namespace":"default"}}
2024-02-23T21:34:43+08:00       INFO    Enter Reconcile {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume", "DemoVolume": {"name":"demovolume-sample","namespace":"default"}, "namespace": "default", "name": "demovolume-sample", "reconcileID": "c92daf9c-4776-4c60-bf1a-e49f2879555a", "spec": {"Name":"demo","Size":10}, "status": {}}
2024-02-23T21:34:43+08:00       INFO    Enter Reconcile {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume", "DemoVolume": {"name":"demovolume-sample","namespace":"default"}, "namespace": "default", "name": "demovolume-sample", "reconcileID": "a441d1a9-a9fc-4c7d-9a99-4a01f34bc561", "req": {"name":"demovolume-sample","namespace":"default"}}
2024-02-23T21:34:43+08:00       INFO    Enter Reconcile {"controller": "demovolume", "controllerGroup": "demo.demo.jonlimpw.io", "controllerKind": "DemoVolume", "DemoVolume": {"name":"demovolume-sample","namespace":"default"}, "namespace": "default", "name": "demovolume-sample", "reconcileID": "a441d1a9-a9fc-4c7d-9a99-4a01f34bc561", "spec": {"Name":"demo","Size":10}, "status": {"Name":"demo"}}
```
```
jonathan.lim@COMP-PW290173CK k8builder-iteration1 % k get demovolume -o yaml
apiVersion: v1
items:
- apiVersion: demo.demo.jonlimpw.io/v1
  kind: DemoVolume
  metadata:
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"demo.demo.jonlimpw.io/v1","kind":"DemoVolume","metadata":{"annotations":{},"labels":{"app.kubernetes.io/created-by":"demo","app.kubernetes.io/instance":"demovolume-sample","app.kubernetes.io/managed-by":"kustomize","app.kubernetes.io/name":"demovolume","app.kubernetes.io/part-of":"demo"},"name":"demovolume-sample","namespace":"default"},"spec":{"Name":"demo","Size":10}}
    creationTimestamp: "2024-02-23T13:33:57Z"
    generation: 1
    labels:
      app.kubernetes.io/created-by: demo
      app.kubernetes.io/instance: demovolume-sample
      app.kubernetes.io/managed-by: kustomize
      app.kubernetes.io/name: demovolume
      app.kubernetes.io/part-of: demo
    name: demovolume-sample
    namespace: default
    resourceVersion: "146252389"
    uid: eec3987c-5599-48ac-bcd7-dafaec462fbe
  spec:
    Name: demo
    Size: 10
  status:
    Name: demo
kind: List
metadata:
  resourceVersion: ""
```
