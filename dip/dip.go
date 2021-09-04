package dip

import (
	"fmt"

	"github.com/mcholismalik/solid-go/dip/after"
)

type Factory struct {
	MessageSocmed after.IMessageSocmed
	MessageEmail  after.IMessageEmail
	User          after.IUser
}

func Run() {
	fmt.Println("Run dip (dependency inversion principle)")

	messageTemplateWebinar := &after.MessageTemplateWebinar{}
	MessageSocmed := &after.MessageSocmed{
		MessageTemplate: messageTemplateWebinar,
	}
	MessageEmail := &after.MessageEmail{
		MessageTemplate: messageTemplateWebinar,
	}
	User := &after.User{}

	factory := Factory{MessageSocmed, MessageEmail, User}
	Execute(&factory)
}

func Execute(f *Factory) {
	messageSocmedPayload := f.MessageSocmed.Create()
	messageEmailPayload := f.MessageEmail.Create()
	fmt.Println()

	senderSocmed := &after.SenderSocmed{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Message:  messageSocmedPayload.Body,
	}
	senderEmail := &after.SenderEmail{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Subject:  messageEmailPayload.Subject,
		Body:     messageEmailPayload.Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
}
