#!/bin/bash
#delete root qdisk for eth3
tc qdisc del dev eth3 root
#attach root qdisc and create the root class for eth3
tc qdisc add dev eth3 root handle 30: cbq bandwidth 100Mbit avpkt 1000
tc class add dev eth3 parent 30:0 classid 30:1 cbq bandwidth 100Mbit rate \
100Mbit allot 1514 weight 10Mbit prio 8 maxburst 20 avpkt 100
#create the 1Mbps class for sales and accounting
tc class add dev eth3 parent 30:1 classid 30:100 cbq bandwidth 100Mbit rate \
1Mbit allot 1514 weight 128Kbit prio 5 maxburst 20 avpkt 1000 bounded
tc qdisc add dev eth3 parent 30:100 sfq quantum 1514b perturb 15
tc filter add dev eth3 parent 30:0 protocol ip prio 5 u32 match ip dst 192.168.1.0/24 flowid 30:100
#delete root qdisc for eth2
tc qdisc del dev eth2 root
#attach root qdisc and create the root class for eth2
tc qdisc add dev eth2 root handle 20: cbq bandwidth 100Mbit rate \
100Mbit allot 1514 weight 10Mbit prio 8 maxburst 20 avpkg 1000
#create the 2Mbps class for all traffic to executivedep.
tc class add dev eth2 parent 20:1 classid 20:10 cbq bandwidth 100Mbit rate \
2Mbit allot 1514 weight 256Kbit prio 5 maxburst 20 avpkt 100 bounded
#the bit torrent and dc++ class - 512Kbps
tc class add dev eth2 parent 20:10 classid 20:100 cbq bandwidth 100Mbit rate \
512Kbit allot 1514 weight 64Kbit prio 5 maxburst 20avpkt 100 bounded
tc qdisc add dev eth2 parent 20:100 sfq quantum 1514b perturb 15
