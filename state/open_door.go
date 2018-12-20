package state

//import "encoding/json"

import (
	"../message"
)


type OpenDoor struct{
	closeDoor *CloseDoor
}

func (self *OpenDoor) Init (closeDoor *CloseDoor) {
	self.closeDoor  = closeDoor
}

func (self *OpenDoor) Do(msg message.Message) (State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	return nil, nil
}

