package network

import (
	"fmt"
	"net/url"
	"github.com/gorilla/websocket"

	"../message"
	"../settings"
)

type WebsocketClient struct {
	Option   *settings.Option

	conn   *websocket.Conn
	// В канал публикуются данные, полученные от сервера
	RecvData  chan message.Message
	// канал для данных, которые требуется отправить на сервер
	SendData  chan message.Message 
}


func (self* WebsocketClient) start() {
	// прием сообщений от сервера
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
	// отправка сообщений на сервер
	go func () {
		for {
			select {
			case msg := <-self.SendData:
				fmt.Println("send data")
				//text := "Hello world"
				//self.conn.WriteMessage(websocket.TextMessage, []byte(text))
				_ = msg
			}
		}
	}()
}

func (self* WebsocketClient) WSOpen() (error) {
	// Открытие вебсокета и инициализация соединения
	u := url.URL{Scheme: "ws", Host: self.Option.ServerAddr, Path: "/"}
	
	fmt.Println("connect: ", u)

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	self.conn = c
	self.RecvData = make(chan message.Message, 5)
	self.SendData = make(chan message.Message, 5)

	// запускаю ридер/врайтер
	self.start()
	
	return err
}