package domophone

// Управление дверью
type Door struct {
	PinNumber int
}

func (self *Door) Init() (error){
	return nil
}

// Открыть дверь
func (self *Door) Open() {

}

// Закрыть дверь
func (self *Door) Close() {

}