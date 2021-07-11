package after

import "fmt"

type ISender interface {
	SendEmail(msg *MessageEmail)
	SendWhatsapp(msg *MessageWhatsapp)
}

type Sender struct{}

func (s *Sender) SendEmail(msg *MessageEmail) bool {
	fmt.Println("Send email")
	fmt.Println("Sender : ", msg.sender)
	fmt.Println("Receiver : ", msg.receiver)
	fmt.Println("Body : ", msg.body)
	return true
}

func (s *Sender) SendWhatsapp(msg *MessageWhatsapp) bool {
	fmt.Println("Send whatsapp")
	fmt.Println("Sender : ", msg.sender)
	fmt.Println("Receiver : ", msg.receiver)
	fmt.Println("Body : ", msg.body)
	return true
}
