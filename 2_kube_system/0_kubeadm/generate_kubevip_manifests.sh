VIP=192.168.206.5 # Should be in the same subnet 192.168.206.0/24, otherwise cannot bypass WireGuard
SUBNET_PREFIX=192.168.206
INTERFACE="$(ip a | grep $SUBNET_PREFIX | awk '{ print $NF }')"
# KVVERSION=$(curl -sL https://api.github.com/repos/kube-vip/kube-vip/releases | jq -r ".[0].name")
KVVERSION=v0.8.7

# alias kube-vip="ctr image pull ghcr.io/kube-vip/kube-vip:$KVVERSION; ctr run --rm --net-host ghcr.io/kube-vip/kube-vip:$KVVERSION vip /kube-vip"
# alias kube-vip="docker run --network host --rm ghcr.io/kube-vip/kube-vip:$KVVERSION"

function kube-vip()
{
    docker run --network host --rm ghcr.io/kube-vip/kube-vip:$KVVERSION "$@"
}

kube-vip manifest pod \
    --interface $INTERFACE \
    --address $VIP \
    --controlplane \
    --arp \
    --leaderElection \
    --enableNodeLabeling
    # --services not included
