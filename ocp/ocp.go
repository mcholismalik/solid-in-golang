package ocp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/ocp/after"
)

func Run() {
	fmt.Println("Run ocp (open closed principle)")

	messageTemplateCompetition := &after.MessageTemplateCompetition{}
	message := &after.Message{
		MessageTemplateCompetition: messageTemplateCompetition,
	}

	messagePayload := message.Create()
	fmt.Println()

	user := &after.User{}
	sender := &after.Sender{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messagePayload,
	}

	sender.SendWhatsapp()
	fmt.Println()
	sender.SendTelegram()
}
