package main

import (
	"fmt"

	"./state"
	"./message"
	"./domophone"
)

func main() {
	// Описывается конечный автомат состояний устройства.
	// состояние ожидание вызова
	waitCall := &state.WaitCall{}
	// состояние звонок
	domophoneCall := &state.DomophoneCall{}
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
	waitCall.Init(domophoneCall)
	// с вызова -> на снятие трубки, либо -> опять на ожидание.
	domophoneCall.Init(answerPhone, waitCall)
	// со снятой трубки -> открытие двери, либо -> положить трбку.
	answerPhone.Init(openDoor, downPhone)
	// с открытой двери -> в закрытие двери
	openDoor.Init(closeDoor)
	// с закрытия двери -> положить трубку
	closeDoor.Init(downPhone)
	// положить трубку -> только на ожидание
	downPhone.Init(waitCall)

	// детектирование изменение GPIO - на предмет вызова
	detectedCall := &domophone.CallDetect{}

	detectedCall.Init(10)

	// Начальное состояние - ожидание вызова
	var currentState state.State = waitCall
	for {
		var msg  message.Message
		select {
			case msg = <- detectedCall.State:
		}
		fmt.Println(">>>")
		currentState, _ = currentState.Do(msg)	
	}	
}