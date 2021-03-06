package state

//import "encoding/json"

import (
	"fmt"
	"../message"
)

type StartCall struct {
	UserNotif chan message.Message

	stateAnswer *UpPhone
	stateWait  *WaitCall
}

func (self *StartCall) Init (answer *UpPhone, wait *WaitCall) {
	self.stateAnswer = answer
	self.stateWait  = wait
}


func (self *StartCall) Do(request message.Message) (State, *message.Communication, error) {
	var state State = self
	fmt.Println("Идет вызов, уведомляем пользователя и ждем действий от него (состояние: StartCall)")	
		
	// Надо уведомить пользователя о вызове
	self.UserNotif <- "Hello"

	// переход обртно, временно для тестов
	state = self.stateWait
	return state, nil, nil
}

