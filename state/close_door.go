package state

import (
	"../message"
	"../omega2/gpio"
)

type CloseDoor struct {
	closing bool
	ControlDoor *gpio.Out

	stateDownPhone  *DownPhone
}

func (self *CloseDoor) Init (downPhone *DownPhone) {
	self.closing = true
	self.stateDownPhone  = downPhone
}

// закрыть дверь
func (self *CloseDoor) Do(request message.Message) (State, *message.Communication, error) {
	self.ControlDoor.LOW()
	return self.stateDownPhone, nil, nil
}

