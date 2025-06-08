/*
	defaulthttp is the default runner and is what will be used

if no runner is specified
*/
package defaulthttp

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/jonny-burkholder/swarm/internal/models"
)

type defaultRunner struct {
	models.Config
	BaseUrl     string
	Headers     map[string]string
	QueryParams url.Values
	Client      *http.Client
}

// TODO: we need config values here
func New(baseUrl string, headers map[string]string, queryParams url.Values, client ...http.Client) *defaultRunner {
	runner := defaultRunner{
		BaseUrl:     baseUrl,
		Headers:     headers,
		QueryParams: queryParams,
	}

	if len(client) > 0 {
		runner.Client = &client[0]
	}

	return &runner
}

func (runner *defaultRunner) Run(collections []models.Collection) error {
	// TODO: make collections run async if async
	// TODO: handle passing in an http client
	// for each collection
	// TODO: I messed up the nesting here, workers should coordinate to
	// do the number of runs, instead of being multiplicative
	for _, collection := range collections {
		// create # workers for # concurrent runs
		resultChan := make(chan []models.Result)
		wg := &sync.WaitGroup{}
		workers := newWorkers(runner.Concurrent, resultChan, wg, runner.Async)
		// have worker do runs until run counter is complete

		// naive implementation to cycle through workers
		go func() {
			for i := range runner.Runs {
				index := i % runner.Concurrent
				workers[index] <- collection.Requests
			}
		}()

		doneChan := make(chan struct{})

		go func() {
			wg.Wait()
			doneChan <- struct{}{}
		}()

		id := 1
	selectloop:
		for {
			select {
			case results := <-resultChan:
				run := models.Run{
					ID:      id,
					Results: results,
				}
				collection.Mu.Lock()
				collection.Runs = append(collection.Runs, run)
				collection.Mu.Unlock()
				id++      // increment id
				wg.Done() // decrement wait group
			case <-doneChan:
				for _, worker := range workers {
					close(worker)
				}
				break selectloop // putting the label to shut the linter up
			}
		}
	}

	return nil
}
