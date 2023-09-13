package schedule

import (
	"fmt"
	"strings"

	"github.com/fwidjaya20/go-framework/contracts/schedule"
)

type Job struct {
	callback func()
	timing   string
}

func NewJob(callback func()) schedule.Job {
	return &Job{
		callback: callback,
	}
}

func (j *Job) SetTiming(expression string) schedule.Job {
	j.timing = expression
	return j
}

func (j *Job) GetTiming() string {
	return j.timing
}

func (j *Job) GetCallback() func() {
	return j.callback
}

func (j *Job) EverySecond() schedule.Job {
	return j.SetTiming("* * * * * *")
}

func (j *Job) EveryTwoSecond() schedule.Job {
	return j.SetTiming("*/2 * * * * *")
}

func (j *Job) EveryThreeSecond() schedule.Job {
	return j.SetTiming("*/3 * * * * *")
}

func (j *Job) EveryFourSecond() schedule.Job {
	return j.SetTiming("*/4 * * * * *")
}

func (j *Job) EveryFiveSecond() schedule.Job {
	return j.SetTiming("*/5 * * * * *")
}

func (j *Job) EveryTenSecond() schedule.Job {
	return j.SetTiming("*/10 * * * * *")
}

func (j *Job) EveryFifteenSecond() schedule.Job {
	return j.SetTiming("*/15 * * * * *")
}

func (j *Job) EveryTwentySecond() schedule.Job {
	return j.SetTiming("*/20 * * * * *")
}

func (j *Job) EveryThirtySecond() schedule.Job {
	return j.SetTiming("*/30 * * * * *")
}

func (j *Job) EveryMinute() schedule.Job {
	return j.SetTiming("* * * * *")
}

func (j *Job) EveryTwoMinute() schedule.Job {
	return j.SetTiming("*/2 * * * *")
}

func (j *Job) EveryThreeMinute() schedule.Job {
	return j.SetTiming("*/3 * * * *")
}

func (j *Job) EveryFourMinute() schedule.Job {
	return j.SetTiming("*/4 * * * *")
}

func (j *Job) EveryFiveMinute() schedule.Job {
	return j.SetTiming("*/5 * * * *")
}

func (j *Job) EveryTenMinute() schedule.Job {
	return j.SetTiming("*/10 * * * *")
}

func (j *Job) EveryFifteenMinute() schedule.Job {
	return j.SetTiming("*/15 * * * *")
}

func (j *Job) EveryThirtyMinute() schedule.Job {
	return j.SetTiming("*/30 * * * *")
}

func (j *Job) Hourly() schedule.Job {
	return j.SetTiming("0 * * * *")
}

func (j *Job) HourlyAt(atMinutes ...string) schedule.Job {
	if len(atMinutes) > 0 {
		return j.SetTiming(fmt.Sprintf("%s * * * *", strings.Join(atMinutes, ",")))
	}
	return j.Hourly()
}

func (j *Job) Daily() schedule.Job {
	return j.SetTiming("0 0 * * *")
}

func (j *Job) Weekly() schedule.Job {
	return j.SetTiming("0 0 * * 0")
}

func (j *Job) Monthly() schedule.Job {
	return j.SetTiming("0 0 1 * *")
}
