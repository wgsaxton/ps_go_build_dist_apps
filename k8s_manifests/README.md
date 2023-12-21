# Install App in K8s Cluster
This folder contains the manifest files needed to install the app in a k8s cluster. May also contain some commands or tips to help get it installed.

## Installing manifests from the control host
```
git clone https://github.com/wgsaxton/ps_go_build_dist_apps.git
cd ps_go_build_dist_apps/k8s_manifests/
```
Apply the manifests in order
```
kubectl apply -f registry_ns_deploy_serv.yaml
kubectl apply -f log_deploy_serv.yaml
kubectl apply -f grading_deploy_serv.yaml
kubectl apply -f portal_deploy_serv.yaml
```
Check how everything is running with

`kubectl get all -n gradebookapp`

## Troubleshooting
Some basic troubleshooting commands

`kubectl logs -n gradebookapp [pod name]`

Use another pod to troubleshoot networking
```
# The pod will tear down on exit
kubectl run -i --tty --rm -n gradebookapp --image=nicolaka/netshoot --restart=Never netshoot -- sh
# Keep the pod running ...
kubectl run -i --tty -n gradebookapp --image=nicolaka/netshoot netshoot -- sh
# then to resume
kubectl attach netshoot -c netshoot -i -t -n gradebookapp
```

