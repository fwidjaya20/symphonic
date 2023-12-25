package jobs

import "github.com/fwidjaya20/symphonic/contracts/schedule"

type Kernel struct{}

func (k *Kernel) Schedule() []schedule.Job {
	return []schedule.Job{
		JobOne(),
		JobTwo(),
	}
}
