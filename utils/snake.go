package utils

import (
	"regexp"
	"strings"
)

func Snake(txt string) string {
	reg := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := reg.ReplaceAllString(txt, "${1}_${2}")
	return strings.ToLower(snake)
}
