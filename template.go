package main

import (
	"encoding/json"
	"html/template"
	"os"
	"strings"
	"strconv"
)

type TemplateStats struct {
	Name        string
	Days        int
	Performance Performance
	Users       Users
	TotalUsers  int
	TotalLines  int
	TotalWords  int
	Events      int
	Speed       float64
	Daily       map[int]int
}

func Produce(c ChanStats) TemplateStats {
	parts := strings.Split(c.channelName, ".")
	var name string
	name = parts[0]
	l, w := LinesAndWords(c)

	return TemplateStats{
		Name:        name,
		Days:        c.stats.impertinent.dayChanges,
		Performance: c.performance,
		Users:       SortedUsers(c, 15),
		TotalUsers:  len(c.stats.relevant.Users),
		Events:      c.stats.impertinent.totalEvents,
		TotalLines:  l,
		Daily:		 c.stats.daily.Lines,
		Speed:       c.speed,
		TotalWords:  w,
	}
}

func LinesAndWords(c ChanStats) (int, int) {
	var lines, words int
	users := c.stats.relevant.Users
	for _, u := range users {
		lines += u.Lines
		words += u.Words
	}
	return lines, words
}

func WriteData(t TemplateStats) {
	dataDir := "output/data/"

	// write top15
	path := dataDir + t.Name
	Top15(path + "_top15.json", t.Users)
	DailyActivity(path + "_daily_activity.json", t.Daily)
}

func Top15(path string, u Users) {
	WriteJSON(path, u)
}

func DailyActivity(path string, d map[int]int) {
	conv := make(map[string]int)
	for k, v := range(d) {
		conv[strconv.Itoa(k)] = v
	}
	WriteJSON(path, conv)
}

func WriteJSON(path string, data interface{}) {
	// write top15
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	// write daily
	defer f.Close()
	d, err := json.Marshal(data)
	f.Write(d)
}

func Output(t TemplateStats) {
	fileName := "output/" + t.Name + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	WriteData(t)

	tpl, err := template.ParseFiles("templates/default.html")
	tplErr := tpl.Execute(file, t)
	if tplErr != nil {
		panic(tplErr)
	}
}
