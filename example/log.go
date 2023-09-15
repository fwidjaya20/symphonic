package main

import (
	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/facades"
	"github.com/fwidjaya20/go-framework/log"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&log.ServiceProvider{},
	})

	facades.App().Boot()

	facades.Logger().Debug("Example Log")
	facades.Logger().Debugf("%s", "Example Log")
}
