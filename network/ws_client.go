package network

import (
	"github.com/gorilla/websocket"
)

type WSClient struct {
	urladdr string
	conn *websocket.Conn
}

func (self* WSClient) WSOpen(urladdr string) (error) {
	self.urladdr = urladdr


	return nil
}