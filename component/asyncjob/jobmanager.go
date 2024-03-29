package asyncjob

import (
	"golang.org/x/net/context"
	"log"
	"sync"
)

type group struct {
	isConcurrent bool
	jobs         []Job
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
	}
}

func (g *group) runJob(ctx context.Context, j Job) error {
	// execute job, with value default
	if err := j.Execute(ctx); err != nil {
		for {
			// print err
			log.Println(err)

			// if err retry failed, break, return error
			if j.State() == StateRetryFailed {
				return err
			}

			// if err retruy succer, return nil
			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}

	return nil
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])

			continue
		}
		job := g.jobs[i]
		errChan <- g.runJob(ctx, job)
		g.wg.Done()
	}

	var err error

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}
	g.wg.Wait()
	return err
}
