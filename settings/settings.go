package settings

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

// Заданы GPIO-порты для управления устройством
const (
	GPIO_CONTROL_PHONE 	= "15"
	GPIO_CONTROL_DOOR 	= "16"
	GPIO_DETECT_CALL   	= "17"
	GPIO_LED_POWER     	= "18"
)

const (
	CONFIG_PATH        	= "/root/config.xml"
)

const (
	SERIAL_PORT		   	= "/dev/ttyS1"
	SERIAL_BOUDRATE	   	= 115200
	SERIAL_BUFSIZE		= 256
)


type Option struct {
	// адрес сервера управления
	ServerAddr	string `xml:"serveraddr"`

	// Логи и пароль пользователя, необходимые для авторизации устройства на сервере
	Login	 	string 	`xml:"login"`
	Password 	string 	`xml:"password"`
}

// Загрузка настроек из конфигурационного файла
func Load() (*Option, error) {
	xmlFile, err := os.Open(CONFIG_PATH)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	data, _ := ioutil.ReadAll(xmlFile)

	option := &Option{}
	if err = xml.Unmarshal(data, option); err != nil {
		return nil, err
	}
	fmt.Printf(">>> %+v \n", option)
	return option, nil
}