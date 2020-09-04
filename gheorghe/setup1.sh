#!/bin/bash
#define where iptables is
IPT=/sbin/iptables
##################Begin the NAT table operations #############
#Flush all the rules in the NAT table
$IPT -t nat -F
#SNAT sales and accounting to port 53 UDP (DNS)
$IPT -t nat -A POSTROUTING -o eth0 -s 192.168.1.0/24 -p --dport 53 -j SNAT --to 1.1.2.96-1.1.2.254
#Transparent Proxy for sales and accounting
$IPT -t nat -A PREROUTING -s 192.168.1.0/24 -p tcp --dport 80 -j REDIRECT --to-port 3128
