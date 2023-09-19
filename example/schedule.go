package main

import (
	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	ExampleKernel "github.com/fwidjaya20/symphonic/example/schedule"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/fwidjaya20/symphonic/schedule"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&schedule.ServiceProvider{},
	})

	facades.App().Boot()

	kernel := ExampleKernel.Kernel{}

	facades.Schedule().Register(kernel.Schedule())

	facades.Schedule().Run()

	select {}
}
