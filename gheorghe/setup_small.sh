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
#SNAT Sales and accounting for HTTPS
$IPT -t nat -A POSTROUTING -o eth0 -s 192.168.1.0/24 -p tcp --dport 443 -j SNAT --to 1.1.2.96-1.1.2.254
#Drop everything else from sales and accounting to the internet
$IPT -t nat -A POSTROUTING -o eth0 -s 192.168.1.0/24 -j DROP
#Transparent Proxy for management
$IPT -t nat -A PREROUTING -s 1.1.2.64/27 -p tcp --dport 80 -j REDIRECT --to-port 3128
#####################END the NAT table operations ##############
#Flush netfilter table
$IPT -F
#allow packets on the loopback interface
$IPT -A INPUT -i lo -J ACCEPT
#delete MANAGEMENT chain if exists
$IPT -X MANAGEMENT
#create MANAGEMENT chain
$IPT -N MANAGEMENT
#add authorized IPs to the MANAGEMENT chain, drop all the others
$IPT -A MANAGEMENT -s 1.1.2.0/26 -j ACCEPT
$IPT -A MANAGEMENT -s 1.1.3.192 -j ACCEPT
$IPT -A MANAGEMENT -s 1.1.9.21 -j ACCEPT
$IPT -A MANAGEMENT -s 1.1.19.61 -j ACCEPT
$IPT -A MANAGEMENT -s 0/0 -j DROP
#Jump incoming packets for port 61146 TCP to the MANAGEMENT chain
$IPT -A INPUT -p tcp --dport 61146 -j MANAGEMENT
#Jump packets destined to 1.1.2.2 port 61146 TCP to the MANAGEMENT chain
$IPT -A FORWARD -d 1.1.2.2 -p tcp --dport 61146 -j MANAGEMENT
#drop samba (netbios and ms-ds)
$IPT -A INPUT -i eth0 -p tcp --dport 137:139 -j DROP
$IPT -A INPUT -i eth0 -p udp --dport 137:139 -j DROP
$IPT -A INPUT -i eth0 -p udp --dport 445 -j DROP
#deny access to the intranet web server
$IPT -A INPUT -i eth0 -p tcp --dport 80 -j DROP
#filter the PostgreSQL port
$IPT -A INPUT -p tcp --dport 5432 -j DROP
#drop incoming TCP SYN packets
$IPT -A INPUT -i eth0 -p tcp --syn -j DROP
#allow http, pop3, smtp for the web and mail server
$IPT -A FORWARD -d 1.1.2.2 -p tcp -m multiport --dport 80,25,110 -j ACCEPT
#drop all other tcp traffic for the web and mail server
$IPT -A FORWARD -d 1.1.2.2 -p tcp --syn -j DROP
#Drop netbios and ms-ds for the managers
$IPT -A FORWARD -d 1.1.2.64/27 -p tcp --dport 137:139 -j DROP
$IPT -A FORWARD -d 1.1.2.64/27 -p udp --dport 137:139 -j DROP
$IPT -A FORWARD -d 1.1.2.64/27 -p tcp --dport 445 -j DROP
#Flush the mangle table
$IPT -t mangle -F
#Mark packets belonging to dc++ and bittorrent
$IPT -t mangle -A POSTROUTING -o eth2 -m layer7 --l7proto bittorrent -j MARK --set-mark 5
$IPT -t mangle -A POSTROUTING -o eth2 -m layer7 --l7proto directconnect -j MARK --set-mark 5

