// Copyright 2012 Antoine Kalmbach <ane@iki.fi>
// Use of this source code is governed by a GPLv2 license
// found in the LICENSE file.
package main

import (
	"fmt"
	"strings"
	"time"
)

type ImpertinentStats struct {
	dayChanges   int
	topicChanges int
	kicks        int
	quits        int
	joins        int
	parts        int
	totalEvents      int
}

type RelevantStats struct {
	Users map[string]UserStats
	Hours HourStats
}
type UserStats struct {
	Lines int `json:"lines"`
	Words int `json:"words"`
}

type HourStats map[int]int

type Performance struct {
	Duration time.Duration
	Cores int
	Threads int
}

type StatsChunk struct {
	impertinent ImpertinentStats
	relevant    RelevantStats
}

type ChanStats struct {
	channelName string
	chunks      []Chunk
	stats       StatsChunk
	performance Performance
	speed       float64
	matcher     interface {
		Matcher
	}
}

func JoinIntMap(a map[int]int, b map[int]int) {
	if len(a) != len(b) {
		return
	}
	for k, v := range a {
		val, present := b[k]
		if present {
			a[k] = v + val
		}
	}
}
func JoinUserStats(a map[string]UserStats, b map[string]UserStats) map[string]UserStats {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	for k, v := range b {
		val, present := a[k]
		if present {
			val.Words += v.Words
			val.Lines += v.Lines
			a[k] = val
		} else {
			a[k] = v
		}
	}
	return a
}

func (a *RelevantStats) Union(b RelevantStats) {
	JoinIntMap(a.Hours, b.Hours)
	a.Users = JoinUserStats(a.Users, b.Users)
}

func (a *ImpertinentStats) Union(b ImpertinentStats) {
	a.dayChanges += b.dayChanges
	a.topicChanges += b.topicChanges
	a.kicks += b.kicks
	a.quits += b.quits
	a.joins += b.joins
	a.parts += b.parts
	a.totalEvents += b.totalEvents
}

func (a *UserStats) Union(b UserStats) {
	a.Lines += b.Lines
	a.Words += b.Words
}

func (a *StatsChunk) Union(b StatsChunk) {
	a.impertinent.Union(b.impertinent)
	a.relevant.Union(b.relevant)
}

func (r RelevantStats) String() string {
	debug.Println(len(r.Users), "users")
	s := make([]string, len(r.Users))
	i := 0
	for nick, v := range r.Users {
		s[i] = fmt.Sprintf("%s: %d lines, %d words", nick, v.Lines, v.Words)
		i++
	}
	return strings.Join(s, "\n")
}

func (s StatsChunk) String() string {
	return fmt.Sprint(s.impertinent) + "\n" + fmt.Sprint(s.relevant)
}
