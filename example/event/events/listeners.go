package events

import (
	"fmt"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
)

type CalculateAuthorStatistic struct{}

func (l *CalculateAuthorStatistic) Event() string {
	return PostCreatedEvent
}

func (l *CalculateAuthorStatistic) Handle(job ContractEvent.Job) error {
	payload, _ := job.GetPayload().(PostCreated)

	fmt.Println("Calculate Author Statistic", l.Event(), payload)

	return nil
}

type SendNewsletterNotification struct{}

func (l SendNewsletterNotification) Event() string {
	return PostCreatedEvent
}

func (l SendNewsletterNotification) Handle(job ContractEvent.Job) error {
	payload, _ := job.GetPayload().(PostCreated)

	fmt.Println("Send Newsletter Notification", l.Event(), payload)

	return nil
}
