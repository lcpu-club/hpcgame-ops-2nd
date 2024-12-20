sudo swapoff -a
sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config

cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v1.32/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v1.32/rpm/repodata/repomd.xml.key
exclude=kubelet kubeadm kubectl cri-tools kubernetes-cni
EOF

sudo dnf install -y kubelet kubeadm kubectl --disableexcludes=kubernetes

sudo systemctl enable --now kubelet

# Manually download as access to docker.com is blocked
wget https://mirrors.pku.edu.cn/docker-ce/linux/rhel/9.4/x86_64/stable/Packages/containerd.io-1.7.24-3.1.el9.x86_64.rpm
dnf install -y ./containerd.io-1.7.24-3.1.el9.x86_64.rpm
rm -f ./containerd.io-1.7.24-3.1.el9.x86_64.rpm

sudo systemctl enable --now containerd

cat <<EOF | sudo tee /etc/crictl.yaml
runtime-endpoint: "unix:///run/containerd/containerd.sock"
timeout: 0
debug: false
EOF

cp -f "$(echo "$0" | dirname)/containerd/config.toml" "/etc/containerd/config.toml"
sudo systemctl restart containerd
