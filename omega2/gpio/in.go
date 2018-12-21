package gpio

import (
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/gpio"
)

// Управление дверью
type In struct {
	PinNumber string

	pin gpio.PinIO
}

func (self *In) Init() error {
	self.pin = gpioreg.ByName(self.PinNumber) // Get GPIO
	//if err := self.pin.In(gpio.PullDown, gpio.RisingEdge); err != nil {
	//if err := self.pin.In(gpio.PullNoChange, gpio.NoEdge); err != nil { // work
	if err := self.pin.In(gpio.Float, gpio.NoEdge); err != nil {
		return err
	}
	return nil
}

// Прочитать значение
func (self *In) Read() (bool){
	var rv = self.pin.Read()
	if rv == true {
		return true
	}
	return false
}


func (self *In) Uinit() {
	self.pin.Halt()
}