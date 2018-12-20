package domophone

// Управление дверью
type ControlDoor struct {
	PinNumber int
}

func (self *ControlDoor) Init() (error){
	return nil
}

// Открыть дверь
func (self *ControlDoor) Open() {

}

// Закрыть дверь
func (self *ControlDoor) Close() {

}