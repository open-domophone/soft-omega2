package message


const (
	TYPE_STATUS_DEVICE = "status"
	TYPE_WEBRTC_COMMUNICATE = "webrtc"

)

// сообщение для передачи/приему по вебсокету
type Communication struct {
	SessionKey string 	`json:"session_key"`
	Type string			`json:"type"`
	Message string		`json:"message"`
}
