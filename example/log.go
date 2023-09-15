package main

import (
	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/foundation"
	"github.com/fwidjaya20/go-framework/log"
)

func main() {
	foundation.App.GetConfig().Add("app.providers", []ContractFoundation.ServiceProvider{
		&log.ServiceProvider{},
	})

	foundation.App.Boot()

	foundation.App.GetLogger().Debug("Example Log")
	foundation.App.GetLogger().Debugf("%s", "Example Log")
}
