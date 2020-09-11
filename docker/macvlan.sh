# on NET1
# MacVLan interface is created similar to all other types on linux 
# network interfaces (using ip command-line tool and link sub command).
sudo ip link add macvlan1 link eth0 type macvlan

# Configure IP 
sudo ip address add 172.16.10.5/24 dev macvlan1

# Bring the interface to UP 
sudo ip link set dev macvlan1 up


