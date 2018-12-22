package state
//
//import "encoding/json"

import (
	"../message"
	"../omega2/gpio"
)

type UpPhone struct{
	ControlPhone *gpio.Out

	stateOpenDoor  *OpenDoor
	stateDownPhone *DownPhone
}

func (self *UpPhone) Init (openDoor *OpenDoor, downPhone *DownPhone) {
	self.stateOpenDoor  = openDoor
	self.stateDownPhone = downPhone
}

func (self *UpPhone) Do(request message.Message) (State, *message.Communication, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	self.ControlPhone.HIGH()
	return nil, nil, nil
}

