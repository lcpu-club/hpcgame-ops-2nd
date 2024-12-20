bash generate_kubevip_manifests.sh | tee /etc/kubernetes/manifests/kube-vip.yaml
# Then configure DNS kubernetes.hpc.lcpu.dev to point to the VIP

VIP=192.168.206.5

kubeadm init --control-plane-endpoint "$VIP:6443" --upload-certs

# Execute the output of the kubeadm init command
# kubeadm join 192.168.0.200:6443 --token 9vr73a.a8uxyaju799qwdjv --discovery-token-ca-cert-hash sha256:7c2e69131a36ae2a042a339b33381c6d0d43887e2de83720eff5359e26aec866 --control-plane --certificate-key f8902e114ef118304e561c3ecd4d0b543adc226b7a07f675f56564185ffe0c07
