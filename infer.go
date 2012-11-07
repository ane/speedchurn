package main

import (
	"github.com/samuel/go-gettext"
	"strings"
)

func ShortDateMonthTranslator(locale string) (func (in string) string) {
	domain, err := gettext.NewDomain("dates", "locale")
	if err != nil { panic("err"); }

	return func(in string) string {
		els := strings.Split(in, " ")
		if len(els) != 2 { panic("invalid date: " + in); }

		day := domain.GetText(locale, els[0])
		month := domain.GetText(locale, els[1])

		return day + " " + month
	}
}
