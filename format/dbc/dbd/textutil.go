package dbd

import (
	"regexp"
	"strings"
)

func nocomma(lyt []string) {
	for i := range lyt {
		lyt[i] = strings.ReplaceAll(lyt[i], ",", "")
	}
}

func match(rgx *regexp.Regexp, str string) bool {
	return rgx.MatchString(str)
}

func search(rgx *regexp.Regexp, str string) []string {
	return rgx.FindStringSubmatch(str)
}
