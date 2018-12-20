package state

import (
	"../message"
	"../domophone"
)

type CloseDoor struct {
	closing bool
	ControlDoor *domophone.ControlDoor

	stateDownPhone  *DownPhone
}

func (self *CloseDoor) Init (downPhone *DownPhone) {
	self.closing = true
	self.stateDownPhone  = downPhone
}

// закрыть дверь
func (self *CloseDoor) Do(msg message.Message) (State, error) {
	self.ControlDoor.Close()
	return self.stateDownPhone, nil
}

