package message

const (
	LINE_WAIT = iota
	LINE_CALL
)

type MessageDomophoneLine struct {
	State int `json:"state"`
}
