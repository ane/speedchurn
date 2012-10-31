package main

type Kick struct {
	Victim []byte
	Kicker []byte
}

type Topic struct {
	Content string
	Changer string
}

type Timestamp struct {
	Hour int
	Minute int
}

type Normal struct {
	Timestamp Timestamp
	Nick string
	Content string
}

type Matcher interface {
	Day([]byte) (bool, bool)
	Kick([]byte) Kick
	Topic([]byte) (bool, Topic)
	Regular([]byte) (bool, Normal)
}

func MatchDayChange(line []byte, m Matcher) (bool, interface{}) {
	match, _ := m.Day(line)
	return match, match
}

func MatchTopicChange(line []byte, m Matcher) (bool, interface{}) {
	return m.Topic(line)
}

func MatchNormal(line []byte, m Matcher) (bool, interface{}) {
	return m.Regular(line)
}
