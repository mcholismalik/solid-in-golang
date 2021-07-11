package before

import "fmt"

type IMessage interface {
	Create()
	Read()
	Send()
}

type Message struct {
	sender   int32
	receiver int32
	body     string
}

func (m *Message) Create() *Message {
	msg := Message{
		sender:   111,
		receiver: 222,
		body:     "Hai rahma !",
	}

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", msg)
	return &msg
}

func (m *Message) Read(msg *Message) *Message {
	fmt.Println("Read Message")
	fmt.Println("Sender : ", m.sender)
	fmt.Println("Receiver : ", m.receiver)
	fmt.Println("Body : ", m.body)
	return msg
}

func (m *Message) Send(msg *Message) *Message {
	fmt.Println("Send Message")
	fmt.Println("Send from : ", m.sender)
	fmt.Println("Receiver to : ", m.receiver)
	fmt.Println("Body is : ", m.body)
	return msg
}
