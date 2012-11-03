// The format module is the "formatting" part of the pipeline, i.e., the part
// where statistics are processed to provide values, e.g., daily mean lines and such.
package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nick  string `json:"nick"`
	Rank  int
	Stats UserStats `json:"stats"`
}

func (a User) String() string {
	return fmt.Sprintf("%s: %d lines, %d words", a.Nick, a.Stats.Lines, a.Stats.Words)
}

type Users []User

func (a Users) Less(i, j int) bool {
	return a[i].Stats.Lines > a[j].Stats.Lines
}

func (a Users) Len() int { return len(a) }

func (a Users) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Take(n int, arr Users) Users {
	i := 0
	var tmp Users
	for ; i < n; i++ {
		a := arr[i]
		a.Rank = i + 1
		tmp = append(tmp, a)
	}
	return tmp
}

func SortedUsers(c ChanStats, limit int) Users {
	users := c.stats.relevant.Users
	var stats Users
	for nick, user := range users {
		stats = append(stats, User{nick, 0, UserStats{user.Lines, user.Words}})
	}
	sort.Sort(stats)
	return Take(15, stats)
}
