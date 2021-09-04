package isp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/isp/after"
)

func Run() {
	fmt.Println("Run isp (interface segregation principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	messageSocmed := &after.MessageSocmed{
		MessageTemplate: messageTemplateWebinar,
	}
	messageEmail := &after.MessageEmail{
		MessageTemplate: messageTemplateWebinar,
	}

	messageSocmedPayload := messageSocmed.Create()
	messageEmailPayload := messageEmail.Create()
	fmt.Println()

	user := &after.User{}
	senderSocmed := &after.SenderSocmed{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messageSocmedPayload.Body,
	}
	senderEmail := &after.SenderEmail{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Subject:  messageEmailPayload.Subject,
		Body:     messageEmailPayload.Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
}
