package after

import "fmt"

type IMessage interface {
	Create() *Message
}

type Message struct {
	senderNumber   int32
	receiverNumber int32
	senderEmail    string
	receiverEmail  string
	body           string
}

func (m *Message) Create() *Message {
	m = &Message{
		senderNumber:   111,
		receiverNumber: 222,
		senderEmail:    "cholis@gmail.com",
		receiverEmail:  "rahma@gmail.com",
		body:           "Hai rahma !",
	}

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return m
}
