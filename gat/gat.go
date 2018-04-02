package main

import (
	"flag"
	"log"
	"os"

	"go-app-tools/gat/cli"
)


func main() {
	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "Show current version.")
	log.Println("gat tool")
	flag.Parse()
	if versionFlag {
		cli.ShowVersion()
		os.Exit(0)
	}
	args := flag.Args()
	err := cli.RunCommand(args[0], args)
	if err != nil {
		log.Fatal(err)
	}
}
