package network

import (
	"fmt"
	"net/url"
	"github.com/gorilla/websocket"

	"../message"
)

type WebsocketClient struct {
	urladdr string
	conn   *websocket.Conn
	// канал для передачи данных внутрь системы
	RecvData  chan message.Message
	// канал в который запишутся данные, которые требуется опубликовать
	SendData  chan message.Message 
}


func (self* WebsocketClient) read() {

	go func() {
		for {
			_, message, err := self.conn.ReadMessage()
			// определить тип сообщения и сконвертировать к соответствующему типу
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			self.RecvData <- "HELLO"
			_ = message
		}

	}()

	for {
		select {
			case msg := <- self.SendData:
				fmt.Println("send data")
				text := "Hello world"
				self.conn.WriteMessage(websocket.TextMessage, []byte(text))
				_ = msg
		}
	}
}

func (self* WebsocketClient) WSOpen(addr string) (error) {
	// Открытие вебсокета и инициализация соединения
	u := url.URL{Scheme: "ws", Host: addr, Path: "/control"}
	
	fmt.Println("connect: ", u)

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	self.conn = c
	self.urladdr = u.String()
	self.RecvData = make(chan message.Message, 5)
	self.SendData = make(chan message.Message, 5)

	// запускаю ридер
	go self.read()
	
	return err
}