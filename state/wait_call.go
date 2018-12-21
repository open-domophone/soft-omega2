package state

import (
	"fmt"
	//"reflect"

	"../message"
)

// Состояние 1. Ожидание вызова домофона
// Висим и слушаем сигнал с оптопары
// При поступлении сигнала всегда идем в состояние "StartCall"
type WaitCall struct{
	stateStartCall *StartCall
}

func (self *WaitCall) Init (call *StartCall) {
	self.stateStartCall = call
}

func (self *WaitCall) Do(msg message.Message) (State, error) {
	var state State = self

	//fmt.Println(reflect.TypeOf(msg))
	if lineMsg, ok := msg.(*message.DomophoneLine); ok {
		if lineMsg.State == message.LINE_CALL {
			fmt.Println("Фиксируем вызов (состояние: WaitCall)")
			//state = self.stateStartCall
		}
	}
	return state, nil
}

