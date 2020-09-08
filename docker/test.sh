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

ip route add 172.16.10.128 via 172.16.10.2
# route can be made persistent by adding below line in interface configuration
auto eth2
iface eth2 net static
address 172.16.10.65
netmask 255.255.255.192
post-up ip route add 172.16.10.192/26 via 172.16.10.66

# delete the configuration of interface
ip address flush dev eth1
ip address flush dev eth2

ip link add host_bridge1

# delete an interface
ip address flush dev eth1
ip address flush dev eth2

ip link add host_bridge1 type bridge
ip link add host_bridge1 type bridge

ip link show

