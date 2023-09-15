package facades

import "github.com/fwidjaya20/go-framework/contracts/schedule"

func Schedule() schedule.Schedule {
	return App().GetSchedule()
}
