package state
//
//import "encoding/json"

import (
	"../message"
	"../domophone"
)

type UpPhone struct{
	ControlPhone *domophone.ControlPhone

	stateOpenDoor  *OpenDoor
	stateDownPhone *DownPhone
}

func (self *UpPhone) Init (openDoor *OpenDoor, downPhone *DownPhone) {
	self.stateOpenDoor  = openDoor
	self.stateDownPhone = downPhone
}

func (self *UpPhone) Do(msg message.Message) (State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	return nil, nil
}

