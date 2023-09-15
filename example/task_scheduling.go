package main

import (
	"fmt"

	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	ContractSchedule "github.com/fwidjaya20/go-framework/contracts/schedule"
	"github.com/fwidjaya20/go-framework/facades"
	"github.com/fwidjaya20/go-framework/schedule"
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
