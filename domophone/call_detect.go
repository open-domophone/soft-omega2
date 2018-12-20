package domophone

import (
	"time"
	"../message"
)

type CallDetect struct {
	PinNumber int
	State 	  chan message.Message
}

func (self *CallDetect) Init(pinNumber int) error {
	self.PinNumber = pinNumber
	self.State     = make(chan message.Message, 5)
	// считывание с GPIO состояния, вызов сигнализируется через оптопару на плате
	go func() {
		for {
			time.Sleep(2 * time.Second)
			msg := &message.DomophoneLine{}
			msg.State = message.LINE_CALL
			self.State <- msg
		}
	}()
	return nil
}


