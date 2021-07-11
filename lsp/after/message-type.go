package after

import "fmt"

// type
type IMessageType interface {
	Create() string
}

type MessageTypeWebinar struct {
	Title  string
	Author string
}

type MessageTypeCompetition struct {
	Name  string
	Level string
}

func (m *MessageTypeWebinar) Create() string {
	m = &MessageTypeWebinar{
		Title:  "Software engineer in silicon valley",
		Author: "Wiliam",
	}
	res := fmt.Sprintf("Webinar %s, dengan pembicara %s akan dimulai", m.Title, m.Author)
	return res
}

func (m *MessageTypeCompetition) Create() string {
	m = &MessageTypeCompetition{
		Name:  "Competitive Programming",
		Level: "Hard",
	}
	res := fmt.Sprintf("Competition %s, level %s akan dimulai", m.Name, m.Level)
	return res
}
