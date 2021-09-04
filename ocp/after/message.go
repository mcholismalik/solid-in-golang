package after

import "fmt"

type IMessage interface {
	Create() *Message
}

type Message struct {
	Body                       string
	MessageTemplateCompetition *MessageTemplateCompetition
	MessageTemplateWebinar     *MessageTemplateWebinar
}

func (m *Message) Create() string {
	var template string
	if m.MessageTemplateCompetition != nil {
		template = m.MessageTemplateCompetition.Create()
	}
	if m.MessageTemplateWebinar != nil {
		template = m.MessageTemplateWebinar.Create()
	}

	m = &Message{
		Body: "Hai malik !" + template,
	}

	res := m.Body

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return res
}
