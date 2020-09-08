ip addr show eth0
ip link show # on NODE1
ifdown eth1 && sudo ifup eth1
ip addr show dev eth1
ip address add 172.16.10.65/26 dev eth2
ip link set eth2 up
=========================================================
#                          NET2                         #
=========================================================
ip link add dummy0 type dummy
ip address add 172.16.10.129/26 dev dummy0
ip link set dummy0 up
=========================================================
#                          NET3                         #
=========================================================
ip link add dummy0 type dummy
ip address add 172.16.10.193/26 dev dummy0
ip link set dummy0 up

=========================================================
sysctl -w net.ipv4.ip_foward=1

ip route show
