package state

import (
	"../message"
)

type State interface {
	Do(msg message.Message) (State, error)
}

