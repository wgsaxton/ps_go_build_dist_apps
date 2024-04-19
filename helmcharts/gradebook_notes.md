# Creating Helm Chart Notes
General notes on how to implement the helm chart

## Initial testing
Initially I just cloned the repo and installed the helm chart that way
```
git clone https://github.com/wgsaxton/ps_go_build_dist_apps.git
cd ps_go_build_dist_apps/helmcharts/
helm install dev gradebook
```
`dev` is the release name I chose

## Pull helm chart from repo
Note that is you specifying a namespace and need it special like with labels, you should create the namespace first. Example, this will add the `istio-injection=enabled` label for istio to work on the namespace.
```
kubectl create namespace gstest
kubectl label namespaces gstest istio-injection=enabled
```
Now install the Helm chart
```
helm upgrade --install dev oci://ghcr.io/wgsaxton/gradebook --version 0.1.0
OR
helm install dev --version 0.1.0 oci://ghcr.io/wgsaxton/gradebook
OR
helm install dev --version 0.1.1 --namespace gstest --create-namespace oci://ghcr.io/wgsaxton/gradebook
OR
helm install dev --version 0.1.2 --namespace gstest oci://ghcr.io/wgsaxton/gradebook
```

So there's no searching the repo, etc. You have to go strait to the helm package and install it using the `helm upgrade/install` cmd noted above. No `helm repo add` option since I'm using github and not a container registry.

## Leaving this at needing to add the Gateway and VirtualService for Istio
I'll watch a few training videos first (Kodekloud)
- Next step - https://istio.io/latest/docs/setup/getting-started/#ip
- https://istio.io/latest/docs/reference/config/networking/gateway/
- https://kubernetes.io/docs/concepts/services-networking/gateway/

## References
- [Store Helm Charts in GitHub Container Repository | Thomas Stringer](https://trstringer.com/helm-charts-github-container-registry/)

