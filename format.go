package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nick string
	UserStats
}

func (a User) String() string {
	return fmt.Sprintf("%s: %d lines, %d words", a.Nick, a.Lines, a.Words)
}

type Users []User

func (a Users) Less(i, j int) bool {
	return a[i].Lines > a[j].Lines
}

func (a Users) Len() int { return len(a) }

func (a Users) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func SortedUsers(c ChanStats) Users {
	users := c.stats.relevant.Users
	var stats Users
	for nick, user := range users {
		stats = append(stats, User{nick, UserStats{user.Lines, user.Words}})
	}
	sort.Sort(stats)
	return stats
}
