package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"tap_code_lai/component/asyncjob"
	"time"
)

func main() {

	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 1 ")

		return errors.New("some thing went wrong at job1")
	})
	//job2 := asyncjob.NewJob(func(ctx context.Context) error {
	//	time.Sleep(time.Second * 2)
	//	log.Println("I am job 2 ")
	//
	//	return nil
	//})
	//
	//job3 := asyncjob.NewJob(func(ctx context.Context) error {
	//	time.Sleep(time.Second)
	//	log.Println("I am job 3 ")
	//
	//	return nil
	//})

	if err := job1.Execute(context.Background()); err != nil {
		log.Println("Job 1 err", err)

		for {
			if err := job1.Retry(context.Background()); err == nil {
				break
			}

			fmt.Println("Job 1 err", err)

			if job1.State() == asyncjob.StateRetryFailed {
				break
			}
		}
	}

}
