package before

import "fmt"

type IMessage interface {
	Create()
	Send()
}

type Message struct {
	body     string
	sender   int
	receiver int
}

func (m *Message) Create() *Message {
	msg := Message{
		body:     "Hi, malik !",
		sender:   6287,
		receiver: 6281,
	}

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", msg)
	return &msg
}

func (m *Message) Send(msg *Message) *Message {
	fmt.Println("Send Message")
	fmt.Println("Sender:", m.sender)
	fmt.Println("Receiver:", m.receiver)
	return msg
}
