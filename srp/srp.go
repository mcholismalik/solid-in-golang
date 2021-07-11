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

	sender := &after.Sender{}
	sender.SendWhatsapp(messagePayload)
	fmt.Println()
	sender.SendEmail(messagePayload)
}
