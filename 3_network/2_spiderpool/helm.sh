helm repo add spiderpool https://spidernet-io.github.io/spiderpool
helm repo update spiderpool
helm install spiderpool spiderpool/spiderpool --namespace kube-system \
    -f "$(echo "$0" | dirname)/spiderpool_values.yaml"
