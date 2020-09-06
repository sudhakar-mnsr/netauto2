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


