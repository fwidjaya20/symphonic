package facades

import (
	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/foundation"
)

func App() ContractFoundation.Application {
	return foundation.App
}
