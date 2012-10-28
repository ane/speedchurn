package main

type ImpertinentStats struct {
	dayChanges int
	topicChanges int
	kicks int
	quits int
	joins int
	parts int
}

type Log struct {
	channelName string
	impertinent ImpertinentStats
}
