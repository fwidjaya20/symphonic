package event

import (
	"fmt"

	ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"
)

type CalculateAuthorStatistic struct{}

func (l *CalculateAuthorStatistic) Handle(e ContractEvent.Job) error {
	payload := e.GetPayload().(PostCreated)
	fmt.Println(l.Signature(), payload)
	return nil
}

func (l *CalculateAuthorStatistic) Signature() string {
	return "Calculate Author Statistic"
}

type SendNewsletterNotification struct{}

func (l SendNewsletterNotification) Handle(e ContractEvent.Job) error {
	payload := e.GetPayload().(PostCreated)
	fmt.Println(l.Signature(), payload)
	return nil
}

func (l SendNewsletterNotification) Signature() string {
	return "Send Newsletter Notification"
}
