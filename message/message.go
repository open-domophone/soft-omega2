package message

const (
	TYPE_SERIAL = iota
	TYPE_CONTROL
	TYPE_VOICE
	TYPE_LINE_DOMOPHONE
)

type Message struct {
	Type int
	Data []byte
}