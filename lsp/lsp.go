package lsp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/lsp/after"
)

func Run() {
	fmt.Println("Run lsp (liskov subtitution principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	message := &after.Message{
		MessageTemplate: messageTemplateWebinar,
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
