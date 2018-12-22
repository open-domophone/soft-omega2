package message

const (
	LINE_WAIT = iota
	LINE_CALL
)

// статус линии домофона - есть вызов или нет
type DomophoneLine struct {
	State int
}
