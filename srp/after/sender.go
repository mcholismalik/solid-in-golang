package after

import "fmt"

type Sender struct {
	Sender   *User
	Receiver *User
	Message  string
}

func (s *Sender) SendWhatsapp() bool {
	fmt.Println("Send whatsapp")
	fmt.Println("Sender:", s.Sender.phone)
	fmt.Println("Receiver:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}

func (s *Sender) SendTelegram() bool {
	fmt.Println("Send telegram")
	fmt.Println("Sender:", s.Sender.phone)
	fmt.Println("Receiver:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}
