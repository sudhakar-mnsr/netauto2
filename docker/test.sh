ip addr show eth0
ip link show # on NODE1
ifdown eth1 && sudo ifup eth1
ip addr show dev eth1
ip address add 172.16.10.65/26 dev eth2
ip link set eth2 up
