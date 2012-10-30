package main

import (
	"bytes"
	"io"
	"reflect"
	"runtime"
	"sync"
)

func ReduceChunks(source chan interface{}, output chan interface{}) {
	accumulated := StatsChunk{}
	for chunk := range source {
		accumulated.Union(chunk.(StatsChunk))
	}
	output <- accumulated
}

func MapChunk(source interface{}, output chan interface{}) {
	chunk := source.(Chunk)
	buffer := bytes.NewBuffer(chunk)
	impStats := ImpertinentStats{}
	for {
		line, err := buffer.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		} else {
			what := Match(line, matcher)
			switch what.(type) {
			default:
				//fmt.Println("type is %T", typ)
			case bool:
				impStats.dayChanges += 1
			case Topic:
				impStats.topicChanges += 1
			}
		}
	}
	output <- StatsChunk{impStats}
}

func Match(line []byte, m Matcher) interface{} {
	methods := []func([]byte, Matcher) (bool, interface{}){MatchDayChange, MatchTopicChange}
	var group sync.WaitGroup
	hit := make(chan interface{})
	// multiplex the matching to all matcher methods
	for _, fun := range methods {
		go func(f func([]byte, Matcher) (bool, interface{})) {
			match, res := f(line, m)
			group.Add(1)
			if match == true {
				hit <- res
			}
			group.Done()
		}(fun)
	}
	go func() {
		group.Wait()
		close(hit)
	}()
	res, ok := <-hit
	if ok {
		return res
	}
	return nil
}

func GetChunks(c ChanStats) chan interface{} {
	chunkChan := make(chan interface{})
	// send chunkchunks one by one
	go func() {
		for _, chunk := range c.chunks {
			chunkChan <- chunk
		}
		close(chunkChan)
	}()
	return chunkChan
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

