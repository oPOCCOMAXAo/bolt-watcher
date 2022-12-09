package utils

import (
	"regexp"
	"strconv"
)

var (
	reNoInt   = regexp.MustCompile(`[^\d]`)
	reNoFloat = regexp.MustCompile(`[^\d\.]`)
)

func TryParseInt(s string) int64 {
	s = reNoInt.ReplaceAllString(s, "")

	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		res = 0
	}

	return res
}

func TryParseFloat(s string) float64 {
	s = reNoFloat.ReplaceAllString(s, "")

	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		res = 0
	}

	return res
}
