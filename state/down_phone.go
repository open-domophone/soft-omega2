package state

//import "encoding/json"

import (
	"../message"
	"../domophone"
)


type DownPhone struct {
	ControlPhone *domophone.ControlPhone

	stateWaitCall  *WaitCall
}

func (self *DownPhone) Init (waitCall *WaitCall) {
	self.stateWaitCall  = waitCall
}

// Опустить трубку и перейти и в ожидание
func (self *DownPhone) Do(msg message.Message) (State, error) {
	var state State = self.stateWaitCall
	self.ControlPhone.Down()
	return state, nil
}

