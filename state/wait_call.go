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

func (self *WaitCall) Do(request message.Message) (State, *message.Communication, error) {
	var state State = self
	var answer *message.Communication = nil

	//fmt.Println(reflect.TypeOf(msg))
	if lineMsg, ok := request.(*message.DomophoneLine); ok {
		if lineMsg.State == message.LINE_CALL {
			// Формируем уведомление о  вызове
			answer = &message.Communication{}
			answer.SessionKey = "session id"
			answer.Type = message.TYPE_STATUS_DEVICE
			answer.Message = "Hello World"
			fmt.Println("Фиксируем вызов (состояние: WaitCall)")
			// переходим на след.состояние
			state = self.stateStartCall
		}
	}
	return state, answer, nil
}

