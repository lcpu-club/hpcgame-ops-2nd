ip link add vxlan0 type vxlan id 342 dstport 9399 dev ens1np0 group 224.100.242.1
ip link set vxlan0 up
ip addr add dev vxlan0 10.234.0.1??/16 # ??: Node ID
ip link set vxlan0 mtu 4450 # 4500 - 50 (vxlan header)
