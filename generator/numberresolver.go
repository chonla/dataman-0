package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

func IntResolver(argsList string) string {
	args := argsSplit(argsList)
	minArg := "0"
	maxArg := "1000000"

	if len(args) > 0 {
		minArg = args[0]
		if len(args) > 1 {
			maxArg = args[1]
		}
	}

	minVal, _ := strconv.ParseInt(minArg, 10, 64)
	maxVal, _ := strconv.ParseInt(maxArg, 10, 64)
	delta := maxVal - minVal
	r := rand.Int63n(delta)

	randNum := fmt.Sprintf("%d", minVal+r)

	return randNum
}
