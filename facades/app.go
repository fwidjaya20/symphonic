package facades

import (
	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	"github.com/fwidjaya20/go-framework/foundation"
)

func App() ContractFoundation.Application {
	return foundation.App
}
