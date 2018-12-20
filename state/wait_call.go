package state

import (
	"fmt"
	"reflect"

	"../message"
)

// Состояние 1. Ожидание вызова домофона
// Висим и слушаем сигнал с оптопары
// При поступлении сигнала всегда идем в состояние "StartCall"
type WaitCall struct{
	call *DomophoneCall
	//call *State
}

func (self *WaitCall) Init (call *DomophoneCall) {
	self.call = call
}

func (self *WaitCall) Do(msg message.Message) (State, error) {
	var state State = self

	//fmt.Println(">>> %T ", msg)
	fmt.Println(reflect.TypeOf(msg))
	if lineMsg, ok := msg.(*message.DomophoneLine); ok {
		if lineMsg.State == message.LINE_CALL {
			fmt.Println("Фиксируем вызов")	
			state = self.call
		}
	}
	return state, nil
}

