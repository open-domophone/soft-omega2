package message

const (
	LINE_WAIT = iota
	LINE_CALL
)

type DomophoneLine struct {
	State int `json:"state"`
}
