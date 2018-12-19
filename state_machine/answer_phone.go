package state_machine
//
//import "encoding/json"

import (
	"../message"
)

// Состояние 3. Поднимаем трубку телефона
type AnswerPhone struct{
	openDoor  *OpenDoor
	downPhone *DownPhone
}

func (self *AnswerPhone) Init (openDoor *OpenDoor, downPhone *DownPhone) {
	self.openDoor  = openDoor
	self.downPhone = downPhone
}

func (self *AnswerPhone) Do(msg* message.Message) (State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	return nil, nil
}

