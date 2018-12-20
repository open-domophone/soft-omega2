package state
//
//import "encoding/json"

import (
	"../message"
)

type AnswerPhone struct{
	openDoor  *OpenDoor
	downPhone *DownPhone
}

func (self *AnswerPhone) Init (openDoor *OpenDoor, downPhone *DownPhone) {
	self.openDoor  = openDoor
	self.downPhone = downPhone
}

func (self *AnswerPhone) Do(msg message.Message) (State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	return nil, nil
}

