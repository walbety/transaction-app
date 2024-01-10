package scheduler

import (
	"github.com/robfig/cron/v3"
)

var (
	runner *cron.Cron
)
const(
	CRON_TIMER = "*/1 * * * *"
)

type Service interface {
	doTask(params string) error
}

type schedulerImpl struct {
	svc Service
}

func Start(service Service) error {
	var err error
	sch := schedulerImpl{
		svc: service,
	}

	runner, err = createJob(CRON_TIMER, sch.trigger)
	if err == nil {
		runner.Start()
	}
	return err
}

func Stop(){
	if runner != nil {
		runner.Stop()
		runner = nil
		// log stop
	}
}

func createJob(timer string, trigger func()) (*cron.Cron, error) {
	// log start

	job := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,)))

	_, err := job.AddFunc(timer, trigger)
	return job, err
}

func (r schedulerImpl) trigger(){
	err := r.svc.doTask("params")
	if err != nil {
		// log error
	}
}