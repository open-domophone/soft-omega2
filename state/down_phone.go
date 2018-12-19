package state

//import "encoding/json"

import (
	"../message"
	"fmt"
)


type DownPhone struct{
	waitCall  *WaitCall
}

func (self *DownPhone) Init (waitCall *WaitCall) {
	self.waitCall  = waitCall
}

func (self *DownPhone) Do(msg* message.Message) (State, error) {
	//if msg.Type == message.TYPE_LINE_DOMOPHONE {
	//	lineMsg := message.MessageDomophoneLine{}
	//	json.Unmarshal(msg.Data, &lineMsg)
	//}
	var state State = self
	fmt.Println("DownPhone - 1")
	return state, nil
	return nil, nil
}

