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
	Day([]byte) (bool, interface{})
	Kick([]byte) Kick
	Topic([]byte) (bool, interface{})
	Regular([]byte) (bool, interface{})
}
