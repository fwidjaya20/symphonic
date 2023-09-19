package main

import (
	"log"
	"time"

	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	"github.com/fwidjaya20/symphonic/event"
	ExampleEvent "github.com/fwidjaya20/symphonic/example/event"
	"github.com/fwidjaya20/symphonic/facades"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&event.ServiceProvider{},
	})

	facades.App().Boot()

	kernel := ExampleEvent.Kernel{}

	facades.Event().Register(kernel.Listen())

	if err := facades.Event().Job(&ExampleEvent.PostCreated{
		Id:        1,
		Author:    "Fredrick Widjaya",
		CreatedAt: time.Now(),
	}).Publish(); nil != err {
		log.Fatalln(err.Error())
	}
}
