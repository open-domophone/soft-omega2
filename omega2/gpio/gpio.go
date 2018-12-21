package gpio

import (
	"periph.io/x/periph/host"
)

func Init() error {
	_, err := host.Init() // Init periph.io
	return err
}