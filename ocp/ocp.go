package ocp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/ocp/after"
)

func Run() {
	fmt.Println("Run ocp (open closed principle)")

	whatsapp := &after.MessageWhatsapp{}
	email := &after.MessageEmail{}
	whatsappPayload := whatsapp.Create()
	emailPayload := email.Create()
	fmt.Println()

	sender := &after.Sender{}
	sender.SendWhatsapp(whatsappPayload)
	fmt.Println()
	sender.SendEmail(emailPayload)
}
