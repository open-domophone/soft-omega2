package omega2

import (
	"fmt"
	"time"
	"runtime"

	"./gpio"
	"../message"
)

type CallDetect struct {
	PinNumber string
	State chan message.Message

	pin gpio.In
}

func (self *CallDetect) Init() error {
	self.State = make(chan message.Message, 5)
	self.pin = gpio.In{PinNumber: self.PinNumber}
	return self.pin.Init()

}


func (self *CallDetect) Run() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			var msg = &message.DomophoneLine{}
			var value = self.pin.Read()
			fmt.Println(">>> StateLine ", value)

			// по умолчанию на gpio подано напряжение от оптопары, при вызове оно упадет,
			// что будет сигнализировать о входящем звонке
			if value == true {
				msg.State = message.LINE_WAIT
			} else {
				msg.State = message.LINE_CALL
			}
			self.State <- msg
			// отдаю управление
			runtime.Gosched()
		}
	}()
}

func (self *CallDetect) Uinit() {
	self.pin.Uinit()
}