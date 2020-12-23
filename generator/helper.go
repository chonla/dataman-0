package generator

import (
	"encoding/csv"
	"strings"
)

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func argsSplit(s string) []string {
	// Split string
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ':'
	fields, err := r.Read()
	if err != nil {
		return []string{s}
	}
	return fields
}
