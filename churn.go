package main

import (
	"time"
	"fmt"
)

// Churn processes a log file first by chunking it into parts and then mapreducing
// all of these simultaneously.
func Churn(file string, ch chan ChanStats) {
	specs := LineChunks(4, file)
	chunks := LoadChunks(file, specs)

	c := ChanStats{channelName: file, chunks: chunks, matcher: new(IrssiMatcher)}

	t := time.Now()
	c.stats = MapReduce(MapChunk, ReduceChunks, GetChunks(c), 4).(StatsChunk)
	dur := time.Since(t)
	fmt.Println(file, "complete in", dur)

	ch <- c
	wg.Done()
}
