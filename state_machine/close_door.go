package state_machine

//import "encoding/json"

import (
	"../message"
)

// Состояние - закрываем дверь - после всегда вешаем трубку
type CloseDoor struct{
	downPhone  *DownPhone
}

func (self *CloseDoor) Init (downPhone *DownPhone) {
	self.downPhone  = downPhone
}

func (self *CloseDoor) Do(msg* message.Message) (*State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	return nil, nil
}

