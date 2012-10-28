package main

type ImpertinentStats struct {
	dayChanges   int
	topicChanges int
	kicks        int
	quits        int
	joins        int
	parts        int
}

type DataChunk struct {
	impertinent ImpertinentStats
}

func (a *ImpertinentStats) Union(b ImpertinentStats) {
	a.dayChanges += b.dayChanges
	a.topicChanges += b.topicChanges
	a.kicks += b.kicks
	a.quits += b.quits
	a.joins += b.joins
	a.parts += b.parts
}

func (a *DataChunk) Union(b DataChunk) {
	a.impertinent.Union(b.impertinent)
}

type ChanStats struct {
	channelName string
	specs       []ChunkSpec
	data DataChunk
}
