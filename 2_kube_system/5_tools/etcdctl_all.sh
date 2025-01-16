#!/bin/bash

pods="$(kubectl get pods -n kube-system -l component=etcd)"

pods=$(echo "$pods" | awk 'NR>1 {print $1}')

for pod in $pods
do

echo "Executing in pod: $pod"
kubectl exec -n kube-system "pod/$pod" -ti -- sh -c "ETCDCTL_API=3 \\
  etcdctl \\
  --endpoints=localhost:2379 \\
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \\
  --cert=/etc/kubernetes/pki/etcd/peer.crt \\
  --key=/etc/kubernetes/pki/etcd/peer.key \\
  \"\$@\"" -- "$@"

done
