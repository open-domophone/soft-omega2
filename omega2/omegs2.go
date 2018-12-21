package omega2

import (
	"os/exec"
	"./gpio"
)

// Инициализация девайса
// настройка GPIO
func InitDevice() error {
	_, err := exec.Command("sh", "-c", "omega2-ctrl gpiomux set spi_s gpio").CombinedOutput()//.Output()
	if err != nil {
		return err
	}
	return gpio.Init()
}
