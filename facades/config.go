package facades

import "github.com/fwidjaya20/go-framework/contracts/config"

func Config() config.Config {
	return App().GetConfig()
}
