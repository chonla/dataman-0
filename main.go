package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chonla/dataman/generator"
	"github.com/chonla/dataman/updater"
	"github.com/mitchellh/go-homedir"
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

	if args[0] == "info" {
		e := printDatasetsInfo()
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
		g, e := generator.New(args[1])

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
	fmt.Printf("%s gen <file.yml>    to generate random date from given <file.yml> config\n", AppName)
	fmt.Printf("%s update            to update internal datasets\n", AppName)
	fmt.Printf("%s info              to show internal datasets stats\n", AppName)
	fmt.Printf("%s version           to show %s version\n", AppName, AppName)
	fmt.Printf("%s help              to show help\n", AppName)
	fmt.Println("")
	fmt.Println("Project page: https://github.com/chonla/dataman")
}

func printVersion() {
	fmt.Printf("%s v.%s", AppName, AppVersion)
}

func updateDatasets() error {
	err := updater.Update()
	return err
}

func printDatasetsInfo() error {
	datasetsPath := "~/.dataman"

	path, err := homedir.Expand(datasetsPath)
	if err != nil {
		return err
	}

	targetPath := fmt.Sprintf("%s/datasets", path)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		fmt.Println("Internal datasets count: 0")
		fmt.Printf("You can use \"%s update\" to update datasets.\n", AppName)
		return nil
	}

	files, err := filepath.Glob(fmt.Sprintf("%s/*.txt", targetPath))
	if err != nil {
		return err
	}

	fmt.Printf("Internal datasets count: %d\n", len(files))

	for _, file := range files {
		_, fileName := filepath.Split(file)
		datasetName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
		fmt.Printf("- %s\n", datasetName)
	}

	fmt.Printf("You can use \"%s update\" to update datasets.\n", AppName)

	return nil
}
