package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Requests) (ParserResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Requests)
	WorkerChan() chan Requests
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Requests)
}

func (e *ConcurrentEngine) Run(seed ...Requests) {
	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.CreatWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seed {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()

		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}

	}

}

func (e *ConcurrentEngine) CreatWorker(in chan Requests, out chan ParserResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}

var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}
