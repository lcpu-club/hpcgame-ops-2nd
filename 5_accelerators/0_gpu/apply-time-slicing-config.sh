kubectl apply -f gpu-time-slicing-config.yaml
kubectl patch clusterpolicies.nvidia.com/cluster-policy \
    -n gpu-operator --type merge \
    -p '{"spec": {"devicePlugin": {"config": {"name": "gpu-time-slicing-config"}}}}'

# kubectl label node <node-name> nvidia.com/device-plugin.config=tesla-t4
# kubectl label node \
#    --selector=nvidia.com/gpu.product=Tesla-T4 \
#    nvidia.com/device-plugin.config=tesla-t4

kubectl label --overwrite node/sc2 nvidia.com/device-plugin.config=v100-16gb-shared
kubectl label --overwrite node/sc4 nvidia.com/device-plugin.config=v100-32gb-shared

kubectl label --overwrite
kubectl label node/l12gpu23 nvidia.com/mig.config=all-2g.20gb --overwrite
kubectl label node/l12gpu25 nvidia.com/mig.config=all-2g.20gb --overwrite
kubectl label node/l12gpu26 nvidia.com/mig.config=all-2g.20gb --overwrite

kubectl label --overwrite node/l12gpu23 nvidia.com/device-plugin.config=a100-80gb-shared
kubectl label --overwrite node/l12gpu11 nvidia.com/device-plugin.config=l40-shared
