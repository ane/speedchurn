// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
package main

type ImpertinentStats struct {
	dayChanges   int
	topicChanges int
	kicks        int
	quits        int
	joins        int
	parts        int
}

type StatsChunk struct {
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

func (a *StatsChunk) Union(b StatsChunk) {
	a.impertinent.Union(b.impertinent)
}

type ChanStats struct {
	channelName string
	chunks      []Chunk
	stats       StatsChunk
	matcher     interface{Matcher}
}
