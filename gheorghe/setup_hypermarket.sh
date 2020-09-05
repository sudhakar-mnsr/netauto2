#Head quarters in one location, one store in same city and
#several stores in other cities.

#hypermarket application uses MS SQL databases in all locations.
#The remote database contains details on stocks and personnel and-
#need to replicate with headquarters database every day at closing.
#Replication is needed for stock details update, as checkout devices-
#query the database for prices and update stocks so that the 
#headquarters database has all info on daily sales and available stocks.

#Software company that also does database administration and remote - 
#storage, so it need access to all databases in every store.

#All locations have IP Analog Telephone Adapters with subscriptions at -
#main provider (HQ). H.323 is the VoIP protocol.

#Headquarters and the store in the same city are connected to same ISP.
#MAN access is much cheaper than an internet connection, headquarters -
#has 10Mbps internet connection with 100Mbps MAN. For the store they -
#wanted only 100 Mbps MAN with no internet connection. 
#The other stores have internet connections from other ISP's in the -
#cities they are in.

####################### BEGIN HEADQUARTERS ######################
#1) Provider assigned public IP 1.1.1.1 for internet connection, this 10Mbps internet connection and 100Mbps metropolitan access
#2) We decided to use private class C 192.168.1.0/24 for internal network.
#3) We set HW router LAN interface with private IP 192.168.1.1
#4) We set MS SQL with static private IP 192.168.1.2
#3) IP ATA must have static private IP 192.168.1.3
####################### END HEADQUARTERS ######################
####################### BEGIN STORE A ######################
#1) The provider assigned private IP address 10.10.12.1 for MAN connection.
#2) This is 100 Mbps metropolitan access and no internet access
#3) We set Linux router A LAN interfaces with private IP 192.168.2.1
#4) We set MS SQL with static private IP 192.168.2.2
#3) IP ATA must have static private IP 192.168.2.3
####################### END STORE A ######################
####################### BEGIN STORE B&C ######################
#1) The provider assigned public IP 1.1.2.1 for B and 1.1.3.1 for C
#2) The connections are 2 Mbps for B and 1 Mbps for C
#3) We use private class 192.168.3.0/24 and 192.168.4.0/24 for B&C resp.
#3) We set Linux routers B&C LAN interfaces with private IPs 192.168.3.1 & 192.168.4.0/24.
#4) We set MS SQL with static private IPs 192.168.3.2(B) & 192.168.4.2(C)
#3) IP ATA must have static private IPs 192.168.3.3 & 192.168.3.4
####################### END STORE B&C ######################
#NOTE: All remote locations must have an encrypted VPN connection to headquarters
#We are using IP tunnel for VPN connection
#Linux distro is Debian
