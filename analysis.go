package main

import (
	"bytes"
	"fmt"
	"io"
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
	for {
		line, err := buffer.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
	output <- StatsChunk{ImpertinentStats{}}
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
