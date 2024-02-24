package main

import (
	"currencyquote/app"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func doJob() {
	hare, err := app.NewHare()
	if err != nil {
		log.Fatalln(err)
		return
	}

	app.UpdateQuotes(*hare)
}

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalln(err)
		return
	}

	j, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(doJob),
	)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	time.Sleep(time.Duration(60 * time.Second))

	sig := make(chan os.Signal, 5)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		log.Fatalln(err)
	}
}
