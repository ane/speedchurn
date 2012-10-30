package main

type Kick struct {
	Victim []byte
	Kicker []byte
}

type Topic struct {
	Content string
	Changer string
}

type Matcher interface {
	Day([]byte) (bool, bool)
	Kick([]byte) Kick
	Topic([]byte) (bool, Topic)
}

func MatchDayChange(line []byte, m Matcher) (bool, interface{}) {
	match, _ := m.Day(line)
	return match, match
}

func MatchTopicChange(line []byte, m Matcher) (bool, interface{}) {
	return m.Topic(line)
}
