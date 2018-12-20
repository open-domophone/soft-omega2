package state

import (
	"../message"
	"../domophone"
)


type OpenDoor struct {
	ControlDoor *domophone.ControlDoor

	stateCloseDoor *CloseDoor
}

func (self *OpenDoor) Init (closeDoor *CloseDoor) {
	self.stateCloseDoor  = closeDoor
}

// открываем дверь и тут-же "закрываем"
func (self *OpenDoor) Do(msg message.Message) (State, error) {
	self.ControlDoor.Open()
	return self.stateCloseDoor, nil
}

