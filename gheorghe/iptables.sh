#!/bin/bash
IP=/sbin/iptables
#... some packet filtering rules
### NAT SECTION
#first of all, we want to flush the NAT table
$IP t nat F
################### SNAT PART
#Jane's special rule.
#Dont SNAT ANY TCP CONNECTIONS from her computer except www and all #udp
#connections except DNS
$IP t nat A POSTROUTING s 192.168.1.31 p tcp -dport ! 80 j DROP
$IP t nat A POSTROUTING s 192.168.1.31 p udp -dport ! 53 j DROP
#Dont SNAT anything from 192.168.1.0/24 to 192.168.2.0/24
$IP t nat A POSTROUTING s 192.168.1.0/24 d 192.168.2.0/24 j ACCEPT
#The boss needs DNAT but we should also SNAT her IP address to 1.2.4.1
$IP t nat A POSTROUTING s 192.168.1.50 j SNAT -to 1.2.4.1
#Snat Everyone
$IP t nat A POSTROUTING s 192.168.1.0/24 o eth1 j SNAT -to 1.2.4.0-1.2.4.32 -to 1.2.3.1
################### DNAT PART
