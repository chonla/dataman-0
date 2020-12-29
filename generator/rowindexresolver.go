package generator

import (
	"fmt"
	"strconv"
)

func RowIndexResolver(argsList string, sessionVars map[string]string) string {
	args := argsSplit(argsList)
	var startFrom int64 = 0

	if len(args) > 0 {
		overridingStartFromArg := args[0]
		startFrom, _ = strconv.ParseInt(overridingStartFromArg, 10, 64)
	}
	var index int64 = 0
	if indexVar, ok := sessionVars["session.index"]; ok {
		index, _ = strconv.ParseInt(indexVar, 10, 64)
	}

	return fmt.Sprintf("%d", index+startFrom)
}
