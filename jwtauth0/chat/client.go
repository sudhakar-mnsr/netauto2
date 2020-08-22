package main
import (
   "github.com/gorilla/websocket"
)

type client struct {
   // socket is the web socket for this client.
   socket *websocket.Conn
   // send is the channel on which messages are sent

