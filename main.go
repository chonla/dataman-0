package main

import (
	"fmt"
	"os"

	"github.com/chonla/dataman/generator"
	"github.com/chonla/dataman/updater"
)

var AppName = "dataman"
var AppVersion = "0.0.0"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	if args[0] == "update" {
		e := updateDatasets()
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if args[0] == "version" {
		printVersion()
		os.Exit(0)
	}

	if args[0] == "help" {
		printUsage()
		os.Exit(0)
	}

	if args[0] == "gen" {
		g, e := generator.New(args[0])

		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}

		g.Generate()
		os.Exit(0)
	}

	printUsage()
	os.Exit(1)
}

func printUsage() {
	fmt.Printf("%s - a random data generator\n", AppName)
	fmt.Println("")
	fmt.Printf("%s gen <file.yml>    generate random date from given <file.yml> config\n", AppName)
	fmt.Printf("%s update            download latest datasets\n", AppName)
	fmt.Printf("%s version           show %s version\n", AppName, AppName)
	fmt.Printf("%s help              show help\n", AppName)
	fmt.Println("")
	fmt.Println("project page: https://github.com/chonla/dataman")
}

func printVersion() {
	fmt.Printf("%s v.%s", AppName, AppVersion)
}

func updateDatasets() error {
	fmt.Println("updating datasets...")
	err := updater.Update()
	return err
}
