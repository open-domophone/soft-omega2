package state

//import "encoding/json"

import (
	"../message"
	"../omega2/gpio"
)


type DownPhone struct {
	ControlPhone *gpio.Out

	stateWaitCall  *WaitCall
}

func (self *DownPhone) Init (waitCall *WaitCall) {
	self.stateWaitCall  = waitCall
}

// Опустить трубку и перейти и в ожидание
func (self *DownPhone) Do(request message.Message) (State, *message.Communication, error) {
	var state State = self.stateWaitCall
	self.ControlPhone.LOW()
	return state, nil, nil
}

