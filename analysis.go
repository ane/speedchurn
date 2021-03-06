package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
	"time"
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
	buffer := bytes.NewBuffer(chunk.Data)
	impStats := ImpertinentStats{}
	relStats := RelevantStats{}
	dayStats := []Day{Day{}}
	relStats.Users = make(map[string]UserStats)
	dayCounter := 0

	translator := ShortDateMonthTranslator(Locale)

	for {
		line, err := buffer.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		} else {
			impStats.totalEvents++
			what := Match(line, matcher)
			switch what.(type) {
			default:
				//fmt.Println("type is %T", typ)
			case []string:
				// TODO: get this into another method, this is awful.

				impStats.dayChanges += 1
				dayCounter++

				day := what.([]string)
				da, month := day[0], day[1]

				// translate first
				translate := fmt.Sprintf("%s %s", da, month)
				translated := translator(translate)

				var transDay, transMonth string
				fmt.Sscanf(translated, "%s %s", &transDay, &transMonth)
				toParse := fmt.Sprintf("%s %s %s", translated, day[2], day[3])

				newDay := Day{Lines: 1}
				parseString := "Mon Jan 02 2006"
				// *both* must differ (i.e. been translated)
				if transDay != da && transMonth != month {
					date, err := time.Parse(parseString, toParse)
					if err != nil {
						panic(err)
					}
					newDay.Date = date
					dayStats = append(dayStats, newDay)
				} else if transDay == da && transMonth == month {
					// try parsing it anyway, maybe it was in english?
					if date, err := time.Parse(parseString, toParse); err == nil {
						newDay.Date = date
					}
					dayStats = append(dayStats, newDay)
				} else {
					// no dice
					dayStats = append(dayStats, newDay)
				}

			case Topic:
				impStats.topicChanges += 1

			case Normal:
				// increment day
				if len(dayStats)-1 <= dayCounter {
					day := dayStats[dayCounter]
					day.Lines++
					dayStats[dayCounter] = day
				}

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
	output <- StatsChunk{impStats, relStats, dayStats}
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
