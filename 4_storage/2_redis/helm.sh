helm repo add ot-helm https://ot-container-kit.github.io/helm-charts/
helm upgrade redis-operator ot-helm/redis-operator \
  --install --create-namespace --namespace ot-operators -f operator_values.yaml

helm upgrade redis-replication ot-helm/redis-replication \
  --install --create-namespace --namespace redis-system -f replication_values.yaml
helm upgrade redis-sentinel ot-helm/redis-sentinel \
  --install --namespace redis-system -f sentinel_values.yaml

# helm uninstall redis-replication -n redis-system
# helm uninstall redis-sentinel -n redis-system
