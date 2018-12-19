package state_machine

import (
	"../message"
)

type State interface {
	Do(msg* message.Message) (State, error)
}

