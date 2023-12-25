package jobs

import (
	"fmt"

	ContractSchedule "github.com/fwidjaya20/symphonic/contracts/schedule"
	"github.com/fwidjaya20/symphonic/schedule"
)

func JobOne() ContractSchedule.Job {
	return schedule.NewJob(func() {
		fmt.Println("Job One")
	}).EverySecond()
}

func JobTwo() ContractSchedule.Job {
	return schedule.NewJob(func() {
		fmt.Println("Job Two")
	}).EveryTwoSecond()
}
