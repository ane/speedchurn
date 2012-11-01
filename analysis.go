package main

import (
	"bytes"
	"io"
	"reflect"
	"runtime"
	"strings"
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
	relStats := RelevantStats{}
	relStats.Users = make(map[string]UserStats)
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
			case Normal:
				w := what.(Normal)
				v, pres := relStats.Users[w.Nick]
				wordCount := len(strings.Split(w.Content, " "))
				if pres {
					v.Lines += 1
					v.Words += wordCount
					relStats.Users[w.Nick] = v
				} else {
					relStats.Users[w.Nick] = UserStats{Lines: 1, Words: wordCount}
				}
			}

		}
	}
	output <- StatsChunk{impStats, relStats}
}

func Match(line []byte, m Matcher) interface{} {
	methods := []func(Matcher, []byte) (bool, interface{}){Matcher.Day, Matcher.Topic, Matcher.Regular}
	// multiplex the matching to all matcher methods
	for i := 0; i < len(methods); i++ {
		match, res := methods[i](m, line)
		if match {
			return res
		}
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
