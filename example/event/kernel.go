package event

import ContractEvent "github.com/fwidjaya20/symphonic/contracts/event"

type Kernel struct{}

func (k *Kernel) Listen() ContractEvent.Collection {
	return ContractEvent.Collection{
		PostCreated{}.Signature(): []ContractEvent.Listener{
			&CalculateAuthorStatistic{},
			&SendNewsletterNotification{},
		},
	}
}
