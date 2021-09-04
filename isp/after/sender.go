package after

import "fmt"

type ISenderSocmed interface {
	SendWhatsapp()
	SendTelegram()
}

type SenderSocmed struct {
	Sender   *User
	Receiver *User
	Message  string
}

type ISenderEmail interface {
	SendEmail()
}

type SenderEmail struct {
	Sender   *User
	Receiver *User
	Subject  string
	Body     string
}

func (s *SenderSocmed) SendWhatsapp() bool {
	fmt.Println("Send whatsapp")
	fmt.Println("SenderEmail:", s.Sender.phone)
	fmt.Println("Receiver:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}

func (s *SenderSocmed) SendTelegram() bool {
	fmt.Println("Send telegram")
	fmt.Println("Sender:", s.Sender.phone)
	fmt.Println("Receiver:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}

func (s *SenderEmail) SendEmail() bool {
	fmt.Println("Send email")
	fmt.Println("Sender:", s.Sender.email)
	fmt.Println("Receiver:", s.Receiver.email)
	fmt.Println("Subject:", s.Subject)
	return true
}
