package domophone


type ControlPhone struct {
	PinNumber int
}

func (self *ControlPhone) Init() (error){
	return nil
}

// Поднять трубку
func (self *ControlPhone) Up() {

}

// Опустить трубку
func (self *ControlPhone) Down() {

}