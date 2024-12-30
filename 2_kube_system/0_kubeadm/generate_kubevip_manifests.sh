VIP=192.168.206.5 # Should be in the same subnet 192.168.206.0/24, otherwise cannot bypass WireGuard
SUBNET_PREFIX=192.168.206
INTERFACE="$(ip a | grep $SUBNET_PREFIX | awk '{ print $NF }')"
# KVVERSION=$(curl -sL https://api.github.com/repos/kube-vip/kube-vip/releases | jq -r ".[0].name")
KVVERSION=v0.8.7

echo "Detected Interface: $INTERFACE" >&2

# alias kube-vip="ctr image pull ghcr.io/kube-vip/kube-vip:$KVVERSION; ctr run --rm --net-host ghcr.io/kube-vip/kube-vip:$KVVERSION vip /kube-vip"
# alias kube-vip="docker run --network host --rm ghcr.io/kube-vip/kube-vip:$KVVERSION"

ctr images pull ghcr.io/kube-vip/kube-vip:$KVVERSION 2>/dev/null 1>/dev/null

function kube-vip()
{
    ctr run --net-host --rm ghcr.io/kube-vip/kube-vip:$KVVERSION "kube-vip-cli-$(openssl rand -hex 4)" /kube-vip "$@"
}

kube-vip manifest pod \
    --interface $INTERFACE \
    --address $VIP \
    --controlplane \
    --arp \
    --leaderElection \
    --enableNodeLabeling \
    | sed "s#path: /etc/kubernetes/super-admin.conf#path: /etc/kubernetes/admin.conf#g" # Fixes: https://www.cnblogs.com/shangmo/p/18441715
    # --services not included
