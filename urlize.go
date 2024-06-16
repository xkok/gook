package gook

import (
	"regexp"
	"strings"
)

type Urlize string

func (u Urlize) Val() string {
	notalnum := regexp.MustCompile(`[^a-z-0-9]+`)
	text := strings.ToLower(string(u))
	text = strings.ReplaceAll(text, " ", "-")
	return string(notalnum.ReplaceAll([]byte(text), []byte("")))
}
