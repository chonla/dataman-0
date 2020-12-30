package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

func IntResolver(argsList string, sessionVars map[string]string) string {
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
	delta := maxVal - minVal + 1
	r := rand.Int63n(delta)

	randNum := fmt.Sprintf("%d", minVal+r)

	return randNum
}

func DecimalResolver(argsList string, sessionVars map[string]string) string {
	args := argsSplit(argsList)
	minArg := "0"
	maxArg := "1000000"
	precisionArg := "5"

	if len(args) > 0 {
		minArg = args[0]
		if len(args) > 1 {
			maxArg = args[1]
			if len(args) > 2 {
				precisionArg = args[2]
			}
		}
	}

	minVal, _ := strconv.ParseInt(minArg, 10, 64)
	maxVal, _ := strconv.ParseInt(maxArg, 10, 64)

	delta := maxVal - minVal + 1
	num := rand.Int63n(delta)

	d := rand.Float64()
	if d > 0.0 && num == maxVal {
		num = num - 1
	}
	r := float64(num) + d

	layout := fmt.Sprintf("%%.%sf", precisionArg)
	randNum := fmt.Sprintf(layout, r)

	return randNum
}
