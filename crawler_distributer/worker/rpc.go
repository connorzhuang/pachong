package worker

import "learngo/crawler/engine"

type CrawlService struct {
}

func (c CrawlService) Process(r Request, result *ParserResult) error {
	engineReq, err := DeserializeRequest(r)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeParserResult(engineResult)
	return nil
}
