package events

import ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"

type Kernel struct{}

func (k *Kernel) Listener() []ContractEvent.Listener {
	return []ContractEvent.Listener{
		&CalculateAuthorStatistic{},
		&SendNewsletterNotification{},
	}
}
