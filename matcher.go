package main

type Kick struct {
	Victim []byte
	Kicker []byte
}

type Topic struct {
	Content []byte
	Changer []byte
	Date []byte
}

type Matcher interface {
	DayChange([]byte) bool
	Kick([]byte) []Kick
	Topic([]byte) []Topic
}

func MatchDayChange(line []byte, m Matcher) (bool) {
	return m.DayChange(line)
}
