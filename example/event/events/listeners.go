package events

import (
	"encoding/json"
	"fmt"
)

type CalculateAuthorStatistic struct{}

func (l *CalculateAuthorStatistic) Handle(payload []byte) error {
	var data PostCreated

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	fmt.Println(l.Signature(), payload)

	return nil
}

func (l *CalculateAuthorStatistic) Signature() string {
	return "Calculate Author Statistic"
}

type SendNewsletterNotification struct{}

func (l SendNewsletterNotification) Handle(payload []byte) error {
	var data PostCreated

	if err := json.Unmarshal(payload, &data); err != nil {
		return err
	}

	fmt.Println(l.Signature(), payload)

	return nil
}

func (l SendNewsletterNotification) Signature() string {
	return "Send Newsletter Notification"
}
