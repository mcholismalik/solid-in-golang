package after

import "fmt"

type ISender interface {
	SendEmail(msg *Message)
	SendWhatsapp(msg *Message)
}

type Sender struct{}

func (s *Sender) SendEmail(msg *Message) bool {
	fmt.Println("Send email")
	fmt.Println("Sender email : ", msg.senderEmail)
	fmt.Println("Receiver email : ", msg.receiverEmail)
	fmt.Println("Body : ", msg.body)
	return true
}

func (s *Sender) SendWhatsapp(msg *Message) bool {
	fmt.Println("Send whatsapp")
	fmt.Println("Sender number : ", msg.senderNumber)
	fmt.Println("Receiver number : ", msg.receiverNumber)
	fmt.Println("Body : ", msg.body)
	return true
}
