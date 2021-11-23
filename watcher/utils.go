package watcher

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`[^\d\.]`)

func parsePrice(s string) float64 {
	s = re.ReplaceAllString(s, "")

	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		res = 0
	}

	return res
}
