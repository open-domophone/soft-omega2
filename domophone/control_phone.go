package domophone


type Phone struct {
	PinNumber int
}

func (self *Phone) Init() (error){
	return nil
}

// Поднять трубку
func (self *Phone) Up() {

}

// Опустить трубку
func (self *Phone) Down() {

}