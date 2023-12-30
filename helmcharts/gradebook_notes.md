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
```
I should add the more detailed way here
this doesn't work yet
```
helm repo add gradebook oci://ghcr.io/wgsaxton/
```