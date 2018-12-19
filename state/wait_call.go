package state
// Состояние 1. Ожидание вызова домофона
// Висим и слушаем сигнал с оптопары
// При поступлении сигнала всегда идем в состояние "StartCall"

import (
	"../message"
	"fmt"
)


type WaitCall struct{
	call *DomophoneCall
	//call *State
}

func (self *WaitCall) Init (call *DomophoneCall) {
	self.call = call
}

func (self *WaitCall) Do(msg* message.Message) (State, error) {
	var state State = self

	fmt.Println("WaitCall - 1")
	state = self.call
	return state, nil
}

