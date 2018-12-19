package main

import (
	"./message"
	state "./state_machine"
)

func main() {
	// Конечный автомат состояний устройства.
	msg := &message.Message{}

	// Ожидание начала вызова
	waitCall := &state.WaitCall{}
	// зарегистрирован завонок
	domophoneCall := &state.DomophoneCall{}
	// Поднять трубку
	answerPhone := &state.AnswerPhone{}
	// Положить Трубку
	downPhone := &state.DownPhone{}
	// Открыть дверь
	openDoor := &state.OpenDoor{}
	// Закрыть дверь
	closeDoor := &state.CloseDoor{}

	// С ожидания можно перейти только на начало вызова либо на само себя
	waitCall.Init(domophoneCall)
	// С состояния вызова можно перейти на снятие трубки или опять на ожидание.
	domophoneCall.Init(answerPhone, waitCall)
	// Снятие трубки и инициализация WebRTC с послед.обменом аудиоданными.
	answerPhone.Init(openDoor, downPhone)
	// Открытие двери -> можно перейти только на закрытие двери
	openDoor.Init(closeDoor)
	// закрытие двери -> можно перейти только на "бросание трубки"
	closeDoor.Init(downPhone)
	// "бросание трубки" -> только на ожидание
	downPhone.Init(waitCall)

	var st state.State
	st = waitCall
	for i := 0; i < 10; i++ {
		st, _ = st.Do(msg)
	}
}