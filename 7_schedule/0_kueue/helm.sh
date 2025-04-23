helm upgrade --install kueue oci://registry.k8s.io/kueue/charts/kueue --version="0.11.3" --create-namespace --namespace=kueue-system -f values.yaml

# helm uninstall kueue --namespace=kueue-system
