# Install rancher-bind backend in rancher
kubectl krew index add rancher-bind https://github.com/Danil-Grigorev/rancher-bind.git
kubectl krew install rancher-bind/rancher-bind
kubectl rancher-bind -f ./example-role.yaml --insecure-skip-tls-verify -d > kubeconfig

# Add kube-bind apiservicebindings to rancher
kubectl krew index add bind https://github.com/kube-bind/krew-index.git
kubectl krew install bind/bind
kubectl bind apiservice --remote-kubeconfig ../rancher-bind/kubeconfig --remote-namespace default -f api.yaml
kubectl bind apiservice --remote-kubeconfig ../rancher-bind/kubeconfig --remote-namespace default -f apitoken.yaml

# In the rancher cluster
kubectl patch apiserviceexport clusterregistrationtokens.management.cattle.io --type merge --patch-file patch.yaml

# Install rancher-turtles in a CAPI management cluster
kubectl apply -f secret.yaml
kind load docker-image --name capi-test ghcr.io/rancher-sandbox/rancher-turtles-amd64:v0.0.1
kind load docker-image --name capi-test registry.k8s.io/capi-operator/cluster-api-operator:v0.5.1
helm install rancher-turtles out/charts/rancher-turtles --set cluster-api-operator.cert-manager.enabled=true --set rancherTurtles.features.embedded-capi.disabled=false --set rancherTurtles.managerArguments[0]=--insecure-skip-verify=true --set cluster-api-operator.cluster-api.configSecret.namespace=default --set cluster-api-operator.cluster-api.configSecret.name=variables --set rancherTurtles.image=ghcr.io/rancher-sandbox/rancher-turtles-amd64 --set rancherTurtles.tag=v0.0.1 --dependency-update -n rancher-turtles-system --create-namespace --wait

kubectl apply -f infra.yaml
kubectl apply -f docker.yaml
