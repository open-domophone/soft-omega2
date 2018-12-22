package main

import (
	"os"
	"os/signal"
	"syscall"
	"runtime"

	"./state"
	"./message"
	"./omega2"
	"./omega2/gpio"
	"./network"
	"./settings"

)

func main() {
	var err error
	var option *settings.Option
	// инициализация omega2
	if err = omega2.InitDevice(); err != nil {
		panic(err)
	}

	if option, err = settings.Load(); err != nil {
		panic(err)
	}

	// GPIO осущ. открытие двери, снятие трубки, индикация питания
	// Порт отвечает за эмитаци. поднятия трубки
	var controlPhone = &gpio.Out{PinNumber: settings.GPIO_CONTROL_PHONE}
	// Порт отвечает за открытие двери
	var controlDoor  = &gpio.Out{PinNumber: settings.GPIO_CONTROL_DOOR}
	// Порт отвечает за индикацию питания (работы программы)
	var ledPower = &gpio.Out{PinNumber: settings.GPIO_LED_POWER}

	// Инициализирую GPIO-порты
	controlPhone.Init()
	controlDoor.Init()
	ledPower.Init()
	defer controlDoor.Uinit()
	defer controlPhone.Uinit()
	defer ledPower.Uinit()
	// зажигаем светодиод - показываем что программа работает
	ledPower.HIGH()

	// детектирование изменение GPIO - на предмет вызова
	var сallDetect = &omega2.CallDetect{PinNumber: settings.GPIO_DETECT_CALL}
	if err = сallDetect.Init(); err != nil {
		panic(err)
	}
	defer сallDetect.Uinit()

	// Подключаемся вебсокетом к серверу
	var websocket = network.WebsocketClient{Option: option}
	if err = websocket.WSOpen(); err != nil {
		panic(err)
	}

	// Описывается конечный автомат состояний устройства.
	// состояние ожидание вызова
	var waitCall = &state.WaitCall{}
	// состояние звонок (в параметрах канал вебсокета - для отправки уведомления пользователю)
	var startCall = &state.StartCall{UserNotif: websocket.SendData}
	// состояние поднять трубку
	var upPhone = &state.UpPhone{ControlPhone: controlPhone}
	// состояние положить трубку
	var downPhone = &state.DownPhone{ControlPhone: controlPhone}
	// состояние открыть дверь
	var openDoor = &state.OpenDoor{ControlDoor: controlDoor}
	// состояние закрыть дверь
	var closeDoor = &state.CloseDoor{ControlDoor: controlDoor}
	// Связывание конечного автомата 
	// с ожидания -> на вызов, либо -> опять на ожидание 
	waitCall.Init(startCall)
	// с вызова -> на снятие трубки, либо -> опять на ожидание.
	startCall.Init(upPhone, waitCall)
	// со снятой трубки -> открытие двери, либо -> положить трбку.
	upPhone.Init(openDoor, downPhone)
	// с открытой двери -> в закрытие двери
	openDoor.Init(closeDoor)
	// с закрытия двери -> положить трубку
	closeDoor.Init(downPhone)
	// положить трубку -> только на ожидание
	downPhone.Init(waitCall)


	// сигнал завершения работы программы
	osSignal := make(chan os.Signal)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGINT)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)


	// Начальное состояние - ожидание вызова
	var currentState state.State = waitCall
	// Запускаю детектирование вызова
	сallDetect.Run()

	var isRunning = true
	for isRunning {
		var msg  message.Message
		select {
			// состояние домофонной линии: есть вызов или нет
			case msg = <- сallDetect.State: {}
			// получение информации от сервера
			case msg = <- websocket.RecvData: {}
			// завершаем работу программы -  по сигналу (из терминала) пользователя
			case <- osSignal:
				isRunning = false
		}
		if msg != nil {
			currentState, _ = currentState.Do(msg)
		}
		runtime.Gosched()
	}
}