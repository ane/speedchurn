package main

import (
	"regexp"
)

type IrssiMatcher struct {
	Matcher
}

var nickChars string = "A-Za-z\\[\\]\\\\`_\\^\\{\\|\\}"
var chanChars string = "^\\r\\n\\0\\s,:"
var timeStampPattern string = "\\d{2}:\\d{2}"
var sepPattern string = "\\s-!-\\s"
var channel string = "([!&#\\+]+[" + chanChars + "]+)"
var nickName string = "([" + nickChars + "]+)"

func (im IrssiMatcher) Day(line []byte) (bool, bool) {
	match, _ := regexp.Match("^--- Day changed", line)
	return match, match
}

func (im IrssiMatcher) Topic(line []byte) (bool, Topic) {
	expr := timeStampPattern + sepPattern + "(" + nickName + ")"
	expr += " changed the topic of " + channel + " to: (.*)"
	rel, _ := regexp.Compile(expr)
	result := rel.FindStringSubmatch(string(line))
	topic := Topic{}
	if result != nil {
		topic.Changer = result[1]
		topic.Content = result[4]
	}
	return result != nil, topic
}
