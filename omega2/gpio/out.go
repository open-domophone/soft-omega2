package gpio

import (
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/gpio"
)

// Управление дверью
type Out struct {
	PinNumber string

	pin gpio.PinIO
}

func (self *Out) Init() {
	self.pin = gpioreg.ByName(self.PinNumber) // Get GPIO
}

// Подать напряжение
func (self *Out) HIGH() {
	self.pin.Out(gpio.High)
}

// Снять напряжение :)
func (self *Out) LOW() {
	self.pin.Out(gpio.Low)
}

func (self *Out) Uinit() {
	self.pin.Out(gpio.Low)
	self.pin.Halt()
}