package events

import ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"

type Kernel struct{}

func (k *Kernel) Listen() ContractEvent.Collection {
	return ContractEvent.Collection{
		//nolint:exhaustruct // ignore due to only retrieve the event topics
		PostCreated{}.Topic(): []ContractEvent.Listener{
			&CalculateAuthorStatistic{},
			&SendNewsletterNotification{},
		},
	}
}
