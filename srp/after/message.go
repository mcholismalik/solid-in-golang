package after

import "fmt"

type IMessage interface {
	Create() *Message
}

type Message struct {
	Body string
}

func (m *Message) Create() string {
	m = &Message{
		Body: "Hai malik !",
	}

	res := m.Body

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return res
}
