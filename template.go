package main

import (
	"encoding/json"
	"html/template"
	"os"
	"strings"
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
	Daily       DailyStats
}

func Produce(c ChanStats) TemplateStats {
	parts := strings.Split(c.channelName, ".")
	var name string
	name = parts[0]
	l, w := LinesAndWords(c)
	c.stats.daily = FixFirst(c.stats.daily)

	return TemplateStats{
		Name:        name,
		Days:        c.stats.impertinent.dayChanges,
		Performance: c.performance,
		Users:       SortedUsers(c, 15),
		TotalUsers:  len(c.stats.relevant.Users),
		Events:      c.stats.impertinent.totalEvents,
		TotalLines:  l,
		Daily:       c.stats.daily,
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

func FixFirst(d DailyStats) DailyStats {
	defer func() {
		if r := recover(); r != nil {
			debug.Println("ERROR: Date modification failed, because time stamps could not be read. Daily activity data will be broken!")
			debug.Println("To fix this, specify the time stamp locale using \"-locale\" command line flag.")
			debug.Println("Error:", r)
		}
	}()
	d[0].Date = d[1].Date.AddDate(0, 0, -1)
	newStats := make([]Day, len(d)/3)
	newStats[0] = d[0]
	for i := 1; (i * 3) < len(d)-3; i++ {
		newStats[i] = d[i*3]
	}
	return newStats
}

func WriteData(t TemplateStats) {
	dataDir := "output/data/"
	if _, err := os.Stat(dataDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dataDir, 0755)
		}
	}

	// write top15
	path := dataDir + t.Name
	Top15(path+"_top15.json", t.Users)
	DailyActivity(path+"_daily_activity.json", t.Daily)
}

func Top15(path string, u Users) {
	WriteJSON(path, u)
}

func DailyActivity(path string, d DailyStats) {
	// Make first date yesterday from second day
	WriteJSON(path, d)
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
	if err != nil {
		panic(err)
	}
	f.Write(d)
	debug.Println("\twrote", path)
}

func Output(t TemplateStats) {
	fileName := "output/" + t.Name + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	WriteData(t)
	debug.Println("\twrote", fileName)

	tpl, err := template.ParseFiles("templates/default.html")
	tplErr := tpl.Execute(file, t)
	if tplErr != nil {
		panic(tplErr)
	}
}
