helm repo add juicefs https://juicedata.github.io/charts/
helm repo update

helm upgrade --install juicefs-csi-driver juicefs/juicefs-csi-driver -n kube-system -f ./juicefs_csi_values.yaml
