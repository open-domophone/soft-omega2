package omega2

import (
	"fmt"
	"runtime"

	"github.com/tarm/serial"
	//"github.com/jacobsa/go-serial/serial"
	"../message"
)


// Srial-интерфейс необходим для взаимодействия с stm32
// stm32 выполняет функции АЦП/ЦАП - получает и подает аудиоданные от/в линию домофона
type SerialPort struct {
	PortName 	string
	BaudRate 	int
	BuffSize 	int

	config 		*serial.Config
	port 		*serial.Port
	//port 		io.ReadWriteCloser
	// В канал публикуются данные, полученные от сервера
	RecvData  	chan message.Message
	// канал для данных, которые требуется отправить на сервер
	SendData  	chan message.Message
}


func (self *SerialPort) start() {
	go func () {
		for {
			buf := make([]byte, self.BuffSize)

			n, err := self.port.Read(buf)
			if err != nil {
				fmt.Println(err)
			}
			buf = buf[:n]
			self.RecvData <- buf
			runtime.Gosched()
		}
	}()

	go func() {
		for {
			select {
			case msg := <-self.SendData:
				if data, ok := msg.([]byte); ok {
					fmt.Println("send data")
					self.port.Write(data)
				}
			}
		}
	}()
}

func (self *SerialPort) Open() error {
	var err error
	self.config = &serial.Config{Name: self.PortName, Baud: self.BaudRate}

	if self.port, err = serial.OpenPort(self.config); err != nil {
		return err
	}
	//options := serial.OpenOptions{
	//	PortName: self.PortName,
	//	BaudRate: self.BaudRate,
	//	//DataBits: 8,
	//	//StopBits: 1,
	//	//MinimumReadSize: 4,
	//}
	//p, err := serial.Open(options)
	//if err != nil {
	//	return err
	//}
	//self.port = p
	return nil
}
