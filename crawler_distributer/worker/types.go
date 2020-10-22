package worker

import (
	"errors"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/Parser"
	"learngo/crawler_distributer/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

func SerializeRequest(r engine.Requests) Request {
	name, args := r.Parser.SerializedParser()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}
func DeserializeRequest(r Request) (engine.Requests, error) {
	parser, err := DeserializeParser(r.Parser)
	if err != nil {
		return engine.Requests{}, err
	}
	return engine.Requests{
		Url:    r.Url,
		Parser: parser,
	}, nil
}
func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(Parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(Parser.ParseCity, config.ParseCity), nil
	case config.ProfileParser:
		if username, ok := p.Args.(string); ok {
			return Parser.NewProfileParser(username), nil
		} else {
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown")
	}
}

type ParserResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeParserResult(r engine.ParserResult) ParserResult {
	result := ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeParserResult(r ParserResult) engine.ParserResult {
	result := engine.ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request:%v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}
