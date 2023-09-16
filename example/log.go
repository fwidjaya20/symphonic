package main

import (
	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/fwidjaya20/symphonic/log"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&log.ServiceProvider{},
	})

	facades.App().Boot()

	facades.Logger().Debug("Example Log")
	facades.Logger().Debugf("%s", "Example Log")
}
