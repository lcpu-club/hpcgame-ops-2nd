helm upgrade --install kueue oci://crmirror.lcpu.dev/us-central1-docker.pkg.dev/k8s-staging-images/charts/kueue --version="v0.10.0" --create-namespace --namespace=kueue-system -f values.yaml

# helm uninstall kueue --namespace=kueue-system
