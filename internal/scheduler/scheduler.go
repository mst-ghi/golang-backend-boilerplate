package scheduler

import (
	"log"

	"github.com/go-co-op/gocron/v2"
)

func InitScheduler() {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		log.Fatalf("Error running scheduler")
	}

	yourScheduler(scheduler)

	scheduler.Start()
}
