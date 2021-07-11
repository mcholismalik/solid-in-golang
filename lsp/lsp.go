package lsp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/lsp/after"
)

func Run() {
	fmt.Println("Run lsp (liskov subtitution principle)")

	var webinar after.IMessageType = &after.MessageTypeWebinar{}
	var competition after.IMessageType = &after.MessageTypeCompetition{}

	whatsapp := &after.MessageWhatsapp{}
	email := &after.MessageEmail{}

	whatsappWebinarPayload := whatsapp.Create(webinar)
	emailWebinarPayload := email.Create(webinar)
	fmt.Println()
	whatsappCompetitionPayload := whatsapp.Create(competition)
	emailCompetitionPayload := email.Create(competition)
	fmt.Println()

	sender := &after.Sender{}
	sender.SendWhatsapp(whatsappWebinarPayload)
	sender.SendWhatsapp(whatsappCompetitionPayload)
	fmt.Println()
	sender.SendEmail(emailWebinarPayload)
	sender.SendEmail(emailCompetitionPayload)
}
