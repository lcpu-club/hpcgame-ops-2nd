helm repo add openebs https://openebs.github.io/openebs
helm repo update

helm install openebs --namespace openebs openebs/openebs --create-namespace --values values.yaml
