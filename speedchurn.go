// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

// Churn processes a log file first by chunking it into parts and then mapreducing
// all of these simultaneously.
func Churn(file string, ch chan ChanStats) {
	cs := lineChunks(4, file)
	c := ChanStats{channelName: file, specs: cs}
	c.data = MapReduce(MapChunk, ReduceChunks, GetChunkSpecs(c), 4).(DataChunk)

	ch <- c
	wg.Done()
}

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("Usage: speedchurn <log1> <log2> ... <logN>")
	}

	logs := args[1:]
	ch := make(chan ChanStats)

	for _, file := range logs {
		wg.Add(1)
		go Churn(file, ch)
	}

	go func() { wg.Wait(); close(ch) }()

	for stats := range ch {
		fmt.Println(stats.data)
	}
}
