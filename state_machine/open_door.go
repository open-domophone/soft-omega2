package state_machine

//import "encoding/json"

import (
	"../message"
)

// Состояние - открываем дверь - после всегда переходим на закрытие двери
type OpenDoor struct{
	closeDoor *CloseDoor
}

func (self *OpenDoor) Init (closeDoor *CloseDoor) {
	self.closeDoor  = closeDoor
}

func (self *OpenDoor) Do(msg* message.Message) (State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	return nil, nil
}

