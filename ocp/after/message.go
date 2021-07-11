package after

import "fmt"

type IMessage interface {
	Create()
}

type Message struct {
	body string
}

type MessageWhatsapp struct {
	Message
	receiver int32
	sender   int32
}

type MessageEmail struct {
	Message
	receiver string
	sender   string
}

func (m *MessageWhatsapp) Create() *MessageWhatsapp {
	m = &MessageWhatsapp{
		Message: Message{
			body: "Hai rahma !",
		},
		sender:   111,
		receiver: 222,
	}

	fmt.Println("Create Message Whatsapp")
	fmt.Printf("%+v\n", m)
	return m
}

func (m *MessageEmail) Create() *MessageEmail {
	m = &MessageEmail{
		Message: Message{
			body: "Hai rahma !",
		},
		sender:   "cholis@gmail.com",
		receiver: "rahma@gmail.com",
	}

	fmt.Println("Create Message Email")
	fmt.Printf("%+v\n", m)
	return m
}
