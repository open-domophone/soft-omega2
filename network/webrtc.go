package network

import (
	"encoding/base64"

	"github.com/pions/webrtc"
	"github.com/pions/webrtc/pkg/ice"

	"../message"
)
// Пример
// https://github.com/pions/webrtc/issues/209


type WebRTC struct {
	// Полученные данные от мобильного приложения
	// (могут быть не только аудиопакеты, но сообщения установки соединения)
	RecvData  	chan message.Message
	// Данные для отправки в мобильное приложение
	// (могут быть не только аудиопакеты, но сообщения установки соединения)
	SendData  	chan message.Message


	// конфигурация webrtc
	config webrtc.RTCConfiguration
	// соединение с мобильным приложением
	peerConnection *webrtc.RTCPeerConnection

	// кодек для удаленного видео
	remoteCodec *webrtc.RTCRtpCodec

	// аудио-стрим с мобильного приложения
	remoteTrack *webrtc.RTCTrack
	// локальный трек аудио
	localOpusTrack *webrtc.RTCTrack
}

func (self *WebRTC) Init(clockrate int) error {
	var err error
	const channels = 1
	// регистрирую только audio кодек Opus (как стандартный для браузера)
	webrtc.RegisterCodec(webrtc.NewRTCRtpOpusCodec(webrtc.DefaultPayloadTypeOpus, uint32(clockrate), channels))

	// Prepare the configuration
	self.config = webrtc.RTCConfiguration{
		IceServers: []webrtc.RTCIceServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	self.peerConnection, err = webrtc.New(self.config)
	if err != nil {
		return err
	}

	// обработчик для подключений
	self.peerConnection.OnICEConnectionStateChange = self.onICEConnectionStateChange
	// вызывается при появлении трека с мобильного приложения
	self.peerConnection.OnTrack = self.ontrack

	// создаю исходящую звуковую дорожку
	self.localOpusTrack, err = self.peerConnection.NewRTCTrack(webrtc.DefaultPayloadTypeOpus,
																	"audio", "label")
	if err != nil {
		return err
	}
	// cоздаю каналы для обменна даннымим внутри приложения
	self.RecvData = make(chan message.Message, 10)
	self.SendData = make(chan message.Message, 10)

	return nil
}

// вызывается при установлении соединения для ice-кандидатов
func (self *WebRTC) onICEConnectionStateChange(connectionState ice.ConnectionState)  {

}


// вызывается когда пришла уадиодорожка с удаленного сервера
func (self *WebRTC) ontrack (track *webrtc.RTCTrack) {
	self.remoteCodec = track.Codec
	self.remoteTrack = track
	for {
		p := <- self.remoteTrack.Packets
		// только сначала требуется раскодировать аудио и запаковать
		// в wav, потом отправить на stm32 по serial
		self.RecvData <- p
	}
}


// вызывается из состояния
// пришли параметры для подключения к мобильному приложения (в base64)
// применить параметры и отправить ответ
// возвращает параметры для локального аудиопотока в base64
func (self *WebRTC) OnOffer (remoteSdpBase64 string)  (string, error) {
	var err error
	// Set the remote SessionDescription
	offer := webrtc.RTCSessionDescription{
		Type: webrtc.RTCSdpTypeOffer,
		Sdp:  string(remoteSdpBase64),
	}
	if err = self.peerConnection.SetRemoteDescription(offer); err != nil {
		return "", err
	}
	// Sets the LocalDescription, and starts our UDP listeners
	answer, err := self.peerConnection.CreateAnswer(nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(answer.Sdp)), nil
}

// передача данных на сервер
func (self *WebRTC) StartAudioSend() {
	go func() {
		data := <- self.SendData
		// написать кодирование данных
		_ = data
		//self.localOpusTrack.Samples <-  encode_data
	}()

}

