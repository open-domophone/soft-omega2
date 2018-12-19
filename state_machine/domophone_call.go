package state_machine

//import "encoding/json"

import (
	"../message"
	"fmt"
)

// Состояние 2. Начало вызова - задача информировать пользователя
type DomophoneCall struct {
	answer *AnswerPhone
	wait  *WaitCall
}

func (self *DomophoneCall) Init (answer *AnswerPhone, wait *WaitCall) {
	self.answer = answer
	self.wait  = wait
}

func (self *DomophoneCall) Do(msg* message.Message) (State, error) {
	var state State = self
	fmt.Println("DomophoneCall - 1")
	return state, nil
}

