package scheduler

import (
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func yourScheduler(scheduler gocron.Scheduler) {
	job, err := scheduler.NewJob(
		gocron.DurationJob(
			time.Minute,
		),
		gocron.NewTask(
			func() {
				log.Println("your job is done")
			},
		),
	)

	if err != nil {
		log.Fatalf("Error running your job")
	}

	log.Println("your job is running:", job.ID())
}
