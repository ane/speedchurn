package main

import (
	"html/template"
	"os"
	"strings"
)

type TemplateStats struct {
	Name string
	Days int
	Performance Performance
}


func Produce(c ChanStats) TemplateStats {
	parts := strings.Split(c.channelName, ".")
	var name string
	name = "#" + parts[0]

	return TemplateStats{
		Name: name,
		Days: c.stats.impertinent.dayChanges,
		Performance: c.performance,
	}
}

func Output(t TemplateStats) {
	fileName := "output/" + t.Name + ".html"
	file, err := os.Create(fileName)
	if err != nil { panic(err) }
	defer file.Close()

	tpl, err := template.ParseFiles("templates/default.html")
	tplErr := tpl.Execute(file, t)
	if tplErr != nil { panic(tplErr) }
}
