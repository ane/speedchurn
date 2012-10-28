package main

func Process(source interface{}, output chan interface{}) {
	//spec := source.(ChunkSpec)
	// do something
	output <- DataChunk{ImpertinentStats{}}
}

func GetSpecs(c ChanStats) chan interface{} {
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
