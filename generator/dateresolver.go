package generator

import (
	"fmt"
	"math/rand"
	"time"
)

// DateResolver - date:layout:from:to
func DateResolver(argsList string, sessionVars map[string]string) string {
	args := argsSplit(argsList)
	minTimeArg := "1970-01-01T00:00:00Z"
	maxTimeArg := "2070-01-01T00:00:00Z"
	formatArg := "2006-01-02T15:04:05Z07:00"

	if len(args) > 0 {
		formatArg = trimQuotes(args[0])
		if len(args) > 1 {
			minTimeArg = args[1]
			if len(args) > 2 {
				maxTimeArg = args[2]
			}
		}
	}

	// time.RFC3339
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	minTime, _ := time.Parse(time.RFC3339, minTimeArg)
	maxTime, _ := time.Parse(time.RFC3339, maxTimeArg)

	delta := maxTime.Sub(minTime)
	seconds := int64(delta.Seconds())
	rand.Seed(time.Now().UnixNano())
	secondsDuration := fmt.Sprintf("%ds", rand.Int63n(seconds))
	duration, _ := time.ParseDuration(secondsDuration)
	sec := minTime.Add(duration)

	return sec.Format(formatArg)
}
