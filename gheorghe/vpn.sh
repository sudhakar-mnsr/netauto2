# Special case for VPN where
# subnets are same, two machines having same IP address communication
#Let us suppose we have 192.168.1.0/24 network and we have database
#installed in 192.168.1.60 on both Headquarters and Remote locations.
#Router connecting HQ is R1 with eth0 configured with 1.2.7.1
#Router connecting Remote location R2, eth0 configured with 1.2.8.1
#R1 VPN IP 10.10.10.1
#R2 VPN IP 10.10.10.2

#The plan is to fake HQ network as 192.168.10.0/24 (database 192.168.10.60)
#Remote network to be faked as 192.168.20.0/24 (database 192.168.20.60)

#On Router 2
iptunnel add vpn1 mode gre remote 1.2.8.1 local 1.2.7.1 key 8132912
ifconfig vpn1 10.10.10.1 pointopoint 10.10.10.2 netmask 255.255.255.252

#On Router 2
iptunnel add vpn1 mode gre remote 1.2.7.1 local 1.2.8.1 key 8132912
ifconfig vpn1 10.10.10.2 pointopoint 10.10.10.1 netmask 255.255.255.252

