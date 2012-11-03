package main

import (
	"html/template"
	"encoding/json"
	"os"
	"strings"
)

type TemplateStats struct {
	Name string
	Days int
	Performance Performance
	Users Users
	TotalUsers int
	TotalLines int
	TotalWords int
	Events int
	Speed float64
}


func Produce(c ChanStats) TemplateStats {
	parts := strings.Split(c.channelName, ".")
	var name string
	name = parts[0]
	l, w := LinesAndWords(c)

	return TemplateStats{
		Name: name,
		Days: c.stats.impertinent.dayChanges,
		Performance: c.performance,
		Users: SortedUsers(c, 15),
		TotalUsers: len(c.stats.relevant.Users),
		Events: c.stats.impertinent.totalEvents,
		TotalLines: l,
		Speed: c.speed,
		TotalWords: w,
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
	top15 := dataDir + t.Name + "_top15.json"
	f, err := os.Create(top15)
	if err != nil { panic(err); }
	defer f.Close()
	d, err :=json.Marshal(t.Users)
	f.Write(d)
}

func Output(t TemplateStats) {
	fileName := "output/" + t.Name + ".html"
	file, err := os.Create(fileName)
	if err != nil { panic(err) }
	defer file.Close()

	WriteData(t)

	tpl, err := template.ParseFiles("templates/default.html")
	tplErr := tpl.Execute(file, t)
	if tplErr != nil { panic(tplErr) }
}
