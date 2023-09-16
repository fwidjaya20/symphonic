package facades

import "github.com/fwidjaya20/symphonic/contracts/config"

func Config() config.Config {
	return App().GetConfig()
}
