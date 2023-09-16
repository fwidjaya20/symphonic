package facades

import "github.com/fwidjaya20/symphonic/contracts/schedule"

func Schedule() schedule.Schedule {
	return App().GetSchedule()
}
