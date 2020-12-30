package generator

import (
	"fmt"
	"strconv"
	"time"
)

// SeqNumberResolver - seqNumber:start_from
func SeqNumberResolver(args []string, sessionVars map[string]string) string {
	var startFrom int64 = 1

	if len(args) > 0 {
		overridingStartFromArg := args[0]
		startFrom, _ = strconv.ParseInt(overridingStartFromArg, 10, 64)
	}
	var index int64 = 0
	if indexVar, ok := sessionVars["session.index"]; ok {
		index, _ = strconv.ParseInt(indexVar, 10, 64)
	}

	return fmt.Sprintf("%d", index+startFrom-1)
}

// SeqDateResolver - seqDate:layout:start_from
func SeqDateResolver(args []string, sessionVars map[string]string) string {
	var startFrom = time.Now()

	formatArg := "2006-01-02T15:04:05Z07:00"

	if len(args) > 0 {
		formatArg = trimQuotes(args[0])
		if len(args) > 1 {
			startFromArg := args[1]
			startFrom, _ = time.Parse(time.RFC3339, startFromArg)
		}
	}

	var index int64 = 0
	if indexVar, ok := sessionVars["session.index"]; ok {
		index, _ = strconv.ParseInt(indexVar, 10, 64)
	}

	dur, _ := time.ParseDuration(fmt.Sprintf("%dh", 24*int(index-1)))

	return startFrom.Add(dur).Format(formatArg)
}
