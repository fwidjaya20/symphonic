package main

import (
	"fmt"

	ContractFoundation "github.com/fwidjaya20/symphonic/contracts/foundation"
	ContractSchedule "github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/fwidjaya20/symphonic/facades"
	"github.com/fwidjaya20/symphonic/schedule"
)

func main() {
	facades.Config().Add("app.providers", []ContractFoundation.ServiceProvider{
		&schedule.ServiceProvider{},
	})

	facades.App().Boot()

	facades.Schedule().Register([]ContractSchedule.Job{
		schedule.NewJob(Job1).EverySecond(),
		schedule.NewJob(Job2).EveryTwoSecond(),
		schedule.NewJob(Job3).EveryThreeSecond(),
	})

	facades.Schedule().Run()

	select {}
}

func Job1() {
	fmt.Println("This is JOB 1")
}

func Job2() {
	fmt.Println("This is JOB 2")
}

func Job3() {
	fmt.Println("This is JOB 3")
}
