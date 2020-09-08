ip addr show eth0
ip link show # on NODE1
ifdown eth1 && sudo ifup eth1
ip addr show dev eth1
