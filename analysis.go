package main

func ReduceChunks(source chan interface{}, output chan interface{}) {
	accumulated := DataChunk{}
	for chunk := range source {
		accumulated.Union(chunk.(DataChunk))
	}
	output <- accumulated
}

func MapChunk(source interface{}, output chan interface{}) {
	// spec := source.(ChunkSpec)
	// do something
	output <- DataChunk{ImpertinentStats{}}
}

func GetChunkSpecs(c ChanStats) chan interface{} {
	specChan := make(chan interface{})
	// send chunkspecs one by one
	go func() {
		for _, spec := range c.specs {
			specChan <- spec
		}
		close(specChan)
	}()
	return specChan
}
