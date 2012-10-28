package main

func ReduceChunks(source chan interface{}, output chan interface{}) {
	accumulated := DataChunk{}
	for chunk := range source {
		accumulated.Union(chunk.(DataChunk))
	}
	output <- accumulated
}
