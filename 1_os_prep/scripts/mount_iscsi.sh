iscsiadm -m discovery -t sendtargets -p 192.168.51.10:3343 -I iser
iscsiadm -m node -T iqn.2025-01.org.lcpu.$(hostname | cut -d "." -f 1):storage -p 192.168.51.10:3343 -l -I iser
mount /dev/sdb /var/lib/containerd/
