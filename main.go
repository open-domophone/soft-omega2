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
	websocket := network.WebsocketClient{}
	err = websocket.WSOpen("localhost:8080")
	if err != nil {
		fmt.Println("websocket:", err)
	}

	// Описывается конечный автомат состояний устройства.
	// состояние ожидание вызова
	waitCall := &state.WaitCall{}
	// состояние звонок
	startCall := &state.StartCall{}
	// состояние поднять трубку
	answerPhone := &state.AnswerPhone{}
	// состояние положить трубку
	downPhone := &state.DownPhone{}
	// состояние открыть дверь
	openDoor := &state.OpenDoor{}
	// состояние закрыть дверь
	closeDoor := &state.CloseDoor{}

	// Связывание конечного автомата 
	// с ожидания -> на вызов, либо -> опять на ожидание 
	waitCall.Init(startCall)
	// с вызова -> на снятие трубки, либо -> опять на ожидание.
	startCall.Init(answerPhone, waitCall)
	// со снятой трубки -> открытие двери, либо -> положить трбку.
	answerPhone.Init(openDoor, downPhone)
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