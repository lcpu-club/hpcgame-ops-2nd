#!/bin/bash

pods="$(kubectl get pods -n kube-system -l component=etcd)"

first_pod=$(echo "$pods" | awk 'NR==2 {print $1}')
echo "Executing in pod: $first_pod"

kubectl exec -n kube-system "pod/$first_pod" -ti -- sh -c "ETCDCTL_API=3 \\
  etcdctl \\
  --endpoints=localhost:2379 \\
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \\
  --cert=/etc/kubernetes/pki/etcd/peer.crt \\
  --key=/etc/kubernetes/pki/etcd/peer.key \\
  \"\$@\"" -- "$@"
