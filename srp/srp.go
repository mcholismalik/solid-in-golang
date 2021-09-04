package srp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/srp/after"
)

func Run() {
	fmt.Println("Run srp (single responsibility principle)")

	message := &after.Message{}
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
