cat <<EOF > /etc/yum.repos.d/lustre-client.repo
[lustre-client]
name=lustre-client
baseurl=https://downloads.whamcloud.com/public/lustre/lustre-2.16.1/el9.4/client
enabled=1
type=rpm
repo_gpgcheck=0
gpgcheck=0
EOF

dnf install -y lustre-client lustre-client-devel kmod-lustre-client-devel kmod-lustre-client
# TODO: change this according to Rocky 9.5's latest kernel
cp -r /lib/modules/5.14.0-427.31.1.el9_4.x86_64/extra/lustre-client /lib/modules/5.14.0-427.37.1.el9_4.x86_64/extra/lustre-client

depmod -a
modprobe lustre
lnetctl net add --if ens1np0 --nid 192.168.50.x@tcp # 可能不需要
lnetctl net del --nid 192.168.206.1??@tcp
# 卡住的话 Ctrl+C 一下再重新执行
mount -t lustre 192.168.51.10@tcp:/lustre /lustre
