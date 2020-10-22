package scheduler

import "learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Requests
}

func (s *SimpleScheduler) WorkerChan() chan engine.Requests {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Requests) {
}

func (s *SimpleScheduler) Run() {

}

func (s *SimpleScheduler) Submit(r engine.Requests) {
	go func() { s.workerChan <- r }()
}
