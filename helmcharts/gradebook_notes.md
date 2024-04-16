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
```
helm upgrade --install dev oci://ghcr.io/wgsaxton/gradebook --version 0.1.0
OR
helm install dev --version 0.1.0 oci://ghcr.io/wgsaxton/gradebook
OR
helm install dev --version 0.1.0 --namespace gstest --create-namespace oci://ghcr.io/wgsaxton/gradebook
```
I should add the more detailed way here but ...

There is no more detailed way since ghcr is not a valid helm chart repository so this doesn't work.
```
helm repo add gradebook oci://ghcr.io/wgsaxton/
```
So there's no searching the repo, etc. You have to go strait to the helm package and install it using the `helm upgrade` cmd noted above.

## References
- [Store Helm Charts in GitHub Container Repository | Thomas Stringer](https://trstringer.com/helm-charts-github-container-registry/)

