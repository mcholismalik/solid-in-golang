package after

import "fmt"

type IMessage interface {
	Create(IMessageType) *MessageWhatsapp
}

type Message struct {
	body string
}

// provider
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

func (m *MessageWhatsapp) Create(mT IMessageType) *MessageWhatsapp {
	additional := mT.Create()
	m = &MessageWhatsapp{
		Message: Message{
			body: "Hai rahma ! " + additional,
		},
		sender:   111,
		receiver: 222,
	}

	fmt.Println("Create Message Whatsapp")
	fmt.Printf("%+v\n", m)
	return m
}

func (m *MessageEmail) Create(mT IMessageType) *MessageEmail {
	additional := mT.Create()
	m = &MessageEmail{
		Message: Message{
			body: "Hai rahma !" + additional,
		},
		sender:   "cholis@gmail.com",
		receiver: "rahma@gmail.com",
	}

	fmt.Println("Create Message Email")
	fmt.Printf("%+v\n", m)
	return m
}
