# on NET1, eth0
# MacVLan interface is created similar to all other types on linux 
# network interfaces (using ip command-line tool and link sub command).
sudo ip link add macvlan1 link eth0 type macvlan

# Configure IP 
sudo ip address add 172.16.10.5/24 dev macvlan1

# Bring the interface to UP 
sudo ip link set dev macvlan1 up

# use ip addr show command to see the details

# on NET1, eth1
sudo ip link add macvlan3 link eth1 type macvlan
sudo ip address add 192.168.10.5/24 dev macvlan3
sudo ip link set dev macvlan3 up

# on NET2, eth1
sudo ip link add macvlan4 link eth1 type macvlan
sudo ip address add 192.168.10.6/24 macvlan4
sudo ip link set dev macvlan4 up
