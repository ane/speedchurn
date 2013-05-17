// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
package main

import (
	"flag"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var matcher Matcher = new(IrssiMatcher)
var Locale string
var Cores int

type debugging bool

const debug debugging = true

func Init() {
	args := os.Args

	// command line Flags
	flag.StringVar(&Locale, "locale", "en_UK", "set log locale (e.g. en_US)")
	flag.IntVar(&Cores, "cpus", 1, "number of logical CPUs to use")
	flag.Parse()

	if len(args) < 2 {
		flag.Usage()
	}
}

func main() {
	Init()
	runtime.GOMAXPROCS(Cores)
	logs := flag.Args()

	ch := make(chan ChanStats)

	for _, file := range logs {
		if _, err := os.Stat(file); err != nil {
			debug.Println("Can't find file ", file)
			continue
		}
		wg.Add(1)
		go Churn(file, ch)
	}

	go func() { wg.Wait(); close(ch) }()

	for stats := range ch {
		debug.Println("Writing output...")
		Output(Produce(stats))
	}
}
