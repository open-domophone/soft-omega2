package state

import (
	"../message"
)

// Интерфейс, для описания состояния
type State interface {
	Do(msg message.Message) (State, error)
}

