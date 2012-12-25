package main

import (
	"fmt"
	"runtime"
	"time"
)

// Churn processes a log file first by chunking it into parts and then mapreducing
// all of these simultaneously.
func Churn(file string, ch chan ChanStats) {
	specs := LineChunks(4, file)
	chunks := LoadChunks(file, specs)

	c := ChanStats{channelName: file, chunks: chunks, matcher: new(IrssiMatcher)}

	t := time.Now()
	fmt.Printf("Using %d cores (%d available)\n", runtime.GOMAXPROCS(2), runtime.NumCPU())
	c.stats = MapReduce(MapChunk, ReduceChunks, GetChunks(c), 4).(StatsChunk)
	dur := time.Since(t)
	fmt.Println(file, "complete in", dur)

	ch <- c
	wg.Done()
}
