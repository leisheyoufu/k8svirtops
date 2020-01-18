# k8svirtops
A repo to test kubebuilder which is a tool to generate k8s operator framework

## Generate code with kubebuilder
kubebuilder init --domain loch.com --license apache2 --owner "loch"
kubebuilder create api --group infra --version v1 --kind VirtualMachine
```
[root@k8sdev k8svirtops]# kubebuilder create api --group infra --version v1 --kind VirtualMachine
Create Resource [y/n]
y
Create Controller [y/n]
y
```
make && make install  // will apply crd in config/bases
kustomize build config/default
kubectl get crd virtualmachines.infra.loch.com -o yaml

## Start controller
make run
kubectl apply -f config/samples/infra_v1_virtualmachine.yaml
```
2020-01-18T17:37:05.949-0500    DEBUG   controller-runtime.controller   Successfully Reconciled {"controller": "virtualmachine", "request": "default/virtualmachine-sample"}
1 1G
```
kubectl get VirtualMachine virtualmachine-sample -o yaml
```
apiVersion: infra.loch.com/v1
kind: VirtualMachine
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"infra.loch.com/v1","kind":"VirtualMachine","metadata":{"annotations":{},"name":"virtualmachine-sample","namespace":"default"},"spec":{"cpu":"1","memory":"1G"}}
  creationTimestamp: "2020-01-18T22:37:05Z"
  generation: 1
  name: virtualmachine-sample
  namespace: default
  resourceVersion: "470470"
  selfLink: /apis/infra.loch.com/v1/namespaces/default/virtualmachines/virtualmachine-sample
  uid: 8a448cfe-a8f0-4d93-9156-8f4dbfcc1e80
spec:
  cpu: "1"
  memory: 1G
status:
  status: Running
```