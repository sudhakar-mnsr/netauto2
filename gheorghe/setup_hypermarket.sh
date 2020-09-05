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

