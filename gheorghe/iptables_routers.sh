#!/bin/bash
I=/sbin/iptables
#The INPUT policy of the core routers should be DROP. 
#We should accept SSH, SNMP, Zebra and BGPd status from network admins.
#DNS packets from our DNS Servers, ICMP packets, packets on lo interface
#TCP packets on 179 (for BGP peers)
#INPUT policy ACCEPT at the end DROP

$I -P INPUT ACCEPT
#Flush and zero chains
$I -F
$I -X
$I -N mang
#Lo + dns + icmp + !syn
$I -A INPUT -i lo -j ACCEPT

$I -A INPUT -p tcp ! --syn -j ACCEPT
$I -A INPUT -p icmp -j ACCEPT
$I -A INPUT -p UDP --sport 53 -s our.dns.server.ourcompany.org -j ACCEPT
$I -A INPUT -p UDP --sport 53 -s our.dns.server2.ourcompany.org -j ACCEPT
#Management chain - ssh, snmp, zebra, BGPD
$I -A INPUT -p TCP --dport 22 -j manag #ssh
$I -A INPUT -p UDP --dport 161 -j manag #snmp
$I -A INPUT -p TCP --dport 2601 -j manag #zebra
$I -A INPUT -p TCP --dport 2605 -j manag #bgpd
#Network Administrators IP address
$I -A manag -s 1.1.40.0/27 -j ACCEPT
$I -A manag -s 1.1.169.o/27
$I -A manag -j DROP

#Our BGP peers
$I -A INPUT -s 1.3.2.1 -p TCP --dport 179 -j ACCEPT #provider-2 peering
$I -A INPUT -s 1.3.1.89 -p TCP --dport 179 -j ACCEPT #provider-2 internet
$I -A INPUT -s 1.1.1.190 -p TCP --dport  179 -j ACCEPT #core-1
$I -A INPUT -s 1.1.1.22 -p TCP --dport  179 -j ACCEPT #core-4
#policy DROP for INPUT chain
$I -p INPUT DROP
