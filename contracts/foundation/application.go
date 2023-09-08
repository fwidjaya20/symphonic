package foundation

import "github.com/fwidjaya20/go-framework/contracts/config"

type Application interface {
	Boot()
	Config() config.Config
}
