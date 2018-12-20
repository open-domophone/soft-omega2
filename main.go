package main

import (
	"fmt"
	"./state"
	"./message"
	"./domophone"
	"./network"
)

func main() {
	var err error
	// Подключаемся вебсокетом к серверу
	var websocket = network.WebsocketClient{}
	err = websocket.WSOpen("localhost:8080")
	if err != nil {
		fmt.Println("websocket:", err)
	}

	var controlPhone = &domophone.ControlPhone{}
	var controlDoor  = &domophone.ControlDoor{}

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

	// детектирование изменение GPIO - на предмет вызова
	сallDetect := &domophone.CallDetect{}
	сallDetect.Init(10)

	// Начальное состояние - ожидание вызова
	var currentState state.State = waitCall
	for {
		var msg  message.Message
		select {
			// состояние домофонной линии: есть вызов или нет
			case msg = <- сallDetect.State:
			// получение информации от сервера
			case msg = <- websocket.RecvData:
		}
		fmt.Println(">>>")
		currentState, _ = currentState.Do(msg)	
	}
}