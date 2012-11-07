// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
package main

import (
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var matcher Matcher = new(IrssiMatcher)

type debugging bool

const debug debugging = true

func main() {
	args := os.Args

	trans := ShortDateMonthTranslator("fi_FI")
	debug.Println(trans("Su Pyl"))

	if len(args) < 2 {
		panic("Usage: speedchurn <log1> <log2> ... <logN>")
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	logs := args[1:]
	ch := make(chan ChanStats)

	for _, file := range logs {
		wg.Add(1)
		go Churn(file, ch)
	}

	go func() { wg.Wait(); close(ch) }()

	for stats := range ch {
		users := SortedUsers(stats, 15)
		for _, user := range users {
			debug.Println(user)
		}
		daily := stats.stats.daily
		daily[0].Date = daily[1].Date.AddDate(0,0,-1)

		for i := 0; i < len(daily); i++ {
			debug.Printf("%dth day %s: %d lines", i, daily[i].Date, daily[i].Lines)
		}
		debug.Println("Writing output...")
		Output(Produce(stats))
	}
}
