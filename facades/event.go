package facades

import (
	"github.com/fwidjaya20/symphonic/contracts/event"
)

func Event() event.Event {
	return App().GetEvent()
}
