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

type MessageTypeCompetitionCustom struct {
	MessageTypeCompetition
	Age   string
	Grade string
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

func (m *MessageTypeCompetitionCustom) Create() string {
	m = &MessageTypeCompetitionCustom{
		MessageTypeCompetition: MessageTypeCompetition{
			Name:  "Competitive Programming",
			Level: "Hard",
		},
		Age:   "17-20",
		Grade: "SMA/SMK",
	}
	res := fmt.Sprintf("Competition %s, level %s untuk usia %s tingkat %s akan dimulai", m.Name, m.Level, m.Age, m.Grade)
	return res
}
