helm repo add rocm https://rocm.github.io/gpu-operator
helm repo update
helm install amd-gpu-operator rocm/gpu-operator-charts \
  --namespace kube-amd-gpu \
  --create-namespace \
  -f values.yaml

kubectl apply -f amdgpu-node-rule.yaml
kubectl apply -f deviceconfig.yaml