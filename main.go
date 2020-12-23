package main

import (
	"fmt"
	"os"

	"github.com/chonla/dataman/generator"
)

var AppName = "dataman"
var AppVersion = "0.0.0"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	if args[0] == "version" {
		printVersion()
		os.Exit(0)
	}

	if args[0] == "help" {
		printUsage()
		os.Exit(0)
	}

	g, e := generator.New(args[0])

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	g.Generate()
}

func printUsage() {
	fmt.Println("dataman <configuration-file.yml>")
}

func printVersion() {
	fmt.Printf("%s v.%s", AppName, AppVersion)
}
