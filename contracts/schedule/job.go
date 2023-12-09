package schedule

type Job interface {
	SetTiming(expression string) Job
	GetTiming() string

	GetCallback() func()

	EverySecond() Job
	EveryTwoSecond() Job
	EveryThreeSecond() Job
	EveryFourSecond() Job
	EveryFiveSecond() Job
	EveryTenSecond() Job
	EveryFifteenSecond() Job
	EveryTwentySecond() Job
	EveryThirtySecond() Job
	EveryMinute() Job
	EveryTwoMinute() Job
	EveryThreeMinute() Job
	EveryFourMinute() Job
	EveryFiveMinute() Job
	EveryTenMinute() Job
	EveryFifteenMinute() Job
	EveryThirtyMinute() Job
	Hourly() Job
	HourlyAt(atMinutes ...string) Job
	Daily() Job
	DailyAt(atHours ...string) Job
	Weekly() Job
	WeeklyAt(atDaysOfWeek ...string) Job
	Monthly() Job
	MonthlyAt(atDaysOfMonth ...string) Job
}
