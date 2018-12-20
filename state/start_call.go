package state

//import "encoding/json"

import (
	"../message"
	"fmt"
)

type StartCall struct {
	answer *AnswerPhone
	wait  *WaitCall
}

func (self *StartCall) Init (answer *AnswerPhone, wait *WaitCall) {
	self.answer = answer
	self.wait  = wait
}

func (self *StartCall) Do(msg message.Message) (State, error) {
	var state State = self
	
	fmt.Println("Идет вызов, уведомляем пользователя и ждем действий от него (состояние: StartCall)")	
		
	// переход обртно, временно для тестов
	state = self.wait
	return state, nil
}

