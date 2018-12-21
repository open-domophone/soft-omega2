package state

import (
	"../message"
	"../omega2/gpio"
)


type OpenDoor struct {
	ControlDoor *gpio.Out

	stateCloseDoor *CloseDoor
}

func (self *OpenDoor) Init (closeDoor *CloseDoor) {
	self.stateCloseDoor  = closeDoor
}

// открываем дверь и тут-же "закрываем"
func (self *OpenDoor) Do(msg message.Message) (State, error) {
	self.ControlDoor.HIGH()
	return self.stateCloseDoor, nil
}

