package scheduler

import (
	"learngo/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Requests
	workerChan  chan chan engine.Requests
}

func (q *QueuedScheduler) WorkerChan() chan engine.Requests {
	return make(chan engine.Requests)
}

func (q *QueuedScheduler) Submit(r engine.Requests) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Requests) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Requests)
	q.workerChan = make(chan chan engine.Requests)
	go func() {
		var requestQ []engine.Requests
		var workerQ []chan engine.Requests
		for {
			var activeRequestChan engine.Requests
			var activeWorkerChan chan engine.Requests
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequestChan = requestQ[0]
				activeWorkerChan = workerQ[0]
			}
			select {
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case activeWorkerChan <- activeRequestChan:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
