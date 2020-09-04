#This script does comparison between packet matches of l7 and ipp2
#For this test we consider three apps: 
#Direct connect, BitTorrent and eDonkey
#
#To test this run below command after running this script
#iptables -l FORWARD -n -v 

#L7 filtering
iptables -I FORWARD -m layer7 --l7proto directconnect
iptables -I FORWARD -m layer7 --l7proto bittorrent
iptables -I FORWARD -m layer7 --l7proto edonkey

