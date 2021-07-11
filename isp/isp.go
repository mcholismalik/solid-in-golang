package isp

import (
	"fmt"

	"github.com/mcholismalik/solid-go/isp/after"
)

func Run() {
	fmt.Println("Run isp (interface segregation principle)")

	var webinar after.IMessageType = &after.MessageTypeWebinar{}
	var competition after.IMessageType = &after.MessageTypeCompetition{}

	var whatsapp after.IMessageWhatsapp = &after.MessageWhatsapp{}
	var email after.IMessageEmail = &after.MessageEmail{}

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
