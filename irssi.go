package main

import "regexp"

type IrssiMatcher struct {
	Matcher
}

func (im IrssiMatcher) DayChange(line []byte) bool {
	match, _ := regexp.Match("^--- Day changed to", line)
	return match
}

