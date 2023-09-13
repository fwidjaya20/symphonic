package main

import (
	"fmt"

	ContractFoundation "github.com/fwidjaya20/go-framework/contracts/foundation"
	ContractSchedule "github.com/fwidjaya20/go-framework/contracts/schedule"
	"github.com/fwidjaya20/go-framework/foundation"
	"github.com/fwidjaya20/go-framework/schedule"
)

func init() {
	foundation.App.GetConfig().Add("app.providers", []ContractFoundation.ServiceProvider{
		&schedule.ServiceProvider{},
	})

	foundation.App.Boot()
}

func main() {
	fmt.Println("Task Scheduling")

	foundation.App.GetSchedule().Register([]ContractSchedule.Job{
		schedule.NewJob(Job1).EverySecond(),
		schedule.NewJob(Job2).EveryTwoSecond(),
		schedule.NewJob(Job3).EveryThreeSecond(),
	})

	foundation.App.GetSchedule().Run()

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
