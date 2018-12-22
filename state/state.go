package state

import (
	"../message"
)

// Интерфейс, для описания состояния
type State interface {
	Do(request message.Message) (State, *message.Communication, error)
}

