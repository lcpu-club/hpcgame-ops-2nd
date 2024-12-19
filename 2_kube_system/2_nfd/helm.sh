export NFD_NS=node-feature-discovery
helm repo add nfd https://kubernetes-sigs.github.io/node-feature-discovery/charts
helm repo update
helm install nfd/node-feature-discovery -f nfd_values.yaml --namespace $NFD_NS --create-namespace --generate-name
