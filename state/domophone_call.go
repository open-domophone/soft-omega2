package state

//import "encoding/json"

import (
	"../message"
	"fmt"
)

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

