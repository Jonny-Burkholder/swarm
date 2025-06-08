package defaulthttp

import (
	"bytes"
	"net/http"
	"net/url"
	"sync"

	"github.com/jonny-burkholder/swarm/internal/models"
)

// TODO: add context for cancel

// newWorkers creates new workers to run the requests from a collection. They can send requests either
// syncronously or asyncronously. The function returns a slice of type chan []models.Request, as each
// worker will have its own channel to receive requests on, so that the caller can trivially keep track
// of the number of runs, and close each channel when the runs are complete
func newWorkers(numWorkers int, resChan chan []models.Result, wg *sync.WaitGroup, async bool, client ...*http.Client) []chan []models.Request {
	res := make([]chan []models.Request, numWorkers)

	// TODO: something with the client

	// slightly uglier than checking in the loop,
	// but I'm guessing more performant, if slightly
	if async {
		for i := range numWorkers {
			c := make(chan []models.Request)
			wrk := asyncWorker{
				requestChan: c,
				resultChan:  resChan,
			}
			res[i] = c

			go wrk.run()
		}
	} else {
		for i := range numWorkers {
			c := make(chan []models.Request)
			wrk := syncWorker{
				requestChan: c,
				resultChan:  resChan,
			}
			res[i] = c

			go wrk.run()
		}
	}

	return res
}

type syncWorker struct {
	wg          *sync.WaitGroup
	requestChan chan []models.Request
	resultChan  chan []models.Result
	client      *http.Client
}

type asyncWorker struct {
	wg          *sync.WaitGroup
	requestChan chan []models.Request
	resultChan  chan []models.Result
	client      *http.Client
}

func (w syncWorker) run() {

	if w.client == nil {
		w.client = http.DefaultClient
	}

	for requests := range w.requestChan {
		results := make([]models.Result, len(requests))
		for i, request := range requests {
			result := models.Result{
				Request: request,
			}

			// parse the url
			reqUrl, err := url.Parse(request.Path)
			if err != nil {
				result.Error = err
				results[i] = result
				continue
			}
			// add the query params to the url
			reqUrl.RawQuery = url.Values(request.QueryParams).Encode()

			// prepare the http request
			req, err := http.NewRequest(request.Method, reqUrl.String(), bytes.NewBuffer(request.Body))
			if err != nil {
				result.Error = err
				results[i] = result
				continue
			}
			for k, v := range request.Headers {
				// use Set() for idempotence
				req.Header.Set(k, v)
			}
			if err = request.Auth.Authenticate(req); err != nil {
				result.Error = err
				results[i] = result
				continue
			}

			// send the request
			res, err := w.client.Do(req)
			if err != nil {
				result.Error = err
				results[i] = result
				continue
			}
			// populate the result
			result.StatusCode = res.StatusCode
			b := []byte{}
			_, err = res.Body.Read(b)
			if err != nil {
				result.Error = err
			}
			result.Body = b
			// call any assertions
			// TODO: implement this. Right now we're not doing anything with the response body

			// add the result to the results slice
			results[i] = result
		}
		// return the results to the resultchan
		w.resultChan <- results
	}
}

func (w asyncWorker) run() {
	// TODO: implement the nightmare of async requests
}
