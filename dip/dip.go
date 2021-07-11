package dip

import (
	"fmt"

	"github.com/mcholismalik/solid-go/dip/after"
)

type Factory struct {
	MessageTypeWebinar     after.IMessageType
	MessageTypeCompetition after.IMessageType
	MessageWhatsapp        after.IMessageWhatsapp
	MessageEmail           after.IMessageEmail
}

func Run() {
	fmt.Println("Run dip (dependency inversion principle)")

	var webinar after.IMessageType = &after.MessageTypeWebinar{}
	var competition after.IMessageType = &after.MessageTypeCompetitionCustom{}

	var whatsapp after.IMessageWhatsapp = &after.MessageWhatsapp{}
	var email after.IMessageEmail = &after.MessageEmail{}

	factory := Factory{webinar, competition, whatsapp, email}
	Execute(&factory)

}

func Execute(f *Factory) {
	whatsappWebinarPayload := f.MessageWhatsapp.Create(f.MessageTypeWebinar)
	emailWebinarPayload := f.MessageEmail.Create(f.MessageTypeWebinar)

	fmt.Println()
	whatsappCompetitionPayload := f.MessageWhatsapp.Create(f.MessageTypeCompetition)
	emailCompetitionPayload := f.MessageEmail.Create(f.MessageTypeCompetition)
	fmt.Println()

	sender := &after.Sender{}
	sender.SendWhatsapp(whatsappWebinarPayload)
	sender.SendWhatsapp(whatsappCompetitionPayload)
	fmt.Println()
	sender.SendEmail(emailWebinarPayload)
	sender.SendEmail(emailCompetitionPayload)
}
