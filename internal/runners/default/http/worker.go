package defaulthttp

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/jonny-burkholder/swarm/internal/models"
)

// TODO: add context for cancel

// newWorkers creates new workers to run the requests from a collection. They can send requests either
// syncronously or asyncronously. The function returns a slice of type chan []models.Request, as each
// worker will have its own channel to receive requests on, so that the caller can trivially keep track
// of the number of runs, and close each channel when the runs are complete
func newWorkers(numWorkers int, resChan chan []models.Result, doneChan chan struct{}, async bool, client ...*http.Client) chan chan []models.Request {
	res := make([]chan []models.Request, numWorkers)

	var workerClient *http.Client
	if len(client) > 0 {
		workerClient = client[0]
	}

	// create a buffered channel that will tell the caller what the next available worker is
	// 1/4 of the number of workers ought to do it
	nextChan := make(chan chan []models.Request, numWorkers/4)

	// slightly uglier than checking in the loop,
	// but I'm guessing more performant, if slightly
	if async {
		for i := range numWorkers {
			c := make(chan []models.Request)
			wrk := asyncWorker{
				requestChan: c,
				resultChan:  resChan,
				nextChan:    nextChan,
				doneChan:    doneChan,
				client:      workerClient,
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
				nextChan:    nextChan,
				doneChan:    doneChan,
				client:      workerClient,
			}
			res[i] = c

			go wrk.run()
		}
	}

	return nextChan
}

type syncWorker struct {
	requestChan chan []models.Request
	resultChan  chan []models.Result
	nextChan    chan chan []models.Request
	doneChan    chan struct{}
	client      *http.Client
}

type asyncWorker struct {
	requestChan chan []models.Request
	resultChan  chan []models.Result
	nextChan    chan chan []models.Request
	doneChan    chan struct{}
	client      *http.Client
}

func (w syncWorker) run() {

	if w.client == nil {
		w.client = http.DefaultClient
	}

	go func() {
		<-w.doneChan
		close(w.requestChan)
	}()

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
		w.nextChan <- w.requestChan
	}
}

func (w asyncWorker) run() {
	// TODO: implement the nightmare of async requests
}
