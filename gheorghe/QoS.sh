#!/bin/bash
#delete root qdisk for eth3
tc qdisc del dev eth3 root
#attach root qdisc and create the root class for eth3
tc qdisc add dev eth3 root handle 30: cbq bandwidth 100Mbit avpkt 1000
tc class add dev eth3 parent 30:0 classid 30:1 cbq bandwidth 100Mbit rate \
100Mbit allot 1514 weight 10Mbit prio 8 maxburst 20 avpkt 100
#create the 1Mbps class for sales and accounting
tc class add dev eth3 parent 30:1 classid 30:100 cbq bandwidth 100Mbit rate\
