package after

import (
	"fmt"
	"time"
)

type IMessageTemplate interface {
	Create() string
}

type MessageTemplate struct {
	StartedAt time.Time
	EndedAt   time.Time
}

type MessageTemplateWebinar struct {
	MessageTemplate
	Title  string
	Author string
}

type MessageTemplateCompetition struct {
	MessageTemplate
	Name  string
	Level string
}

func (m *MessageTemplateWebinar) Create() string {
	m = &MessageTemplateWebinar{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Title:  "Software engineer in silicon valley",
		Author: "Wiliam",
	}
	res := fmt.Sprintf("Webinar %s, dengan pembicara %s akan dimulai pada %s sampai %s", m.Title, m.Author, m.StartedAt, m.EndedAt)
	return res
}

func (m *MessageTemplateCompetition) Create() string {
	m = &MessageTemplateCompetition{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Name:  "Competitive Programming",
		Level: "Hard",
	}
	res := fmt.Sprintf("Competition %s, level %s akan dimulai pada %s sampai %s", m.Name, m.Level, m.StartedAt, m.EndedAt)
	return res
}
