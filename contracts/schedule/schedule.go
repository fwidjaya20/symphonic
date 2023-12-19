package schedule

type Schedule interface {
	Register(jobs []Job)
	Run()
	Stop()
}
