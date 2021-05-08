package main

import (
	"flag"
	"fmt"
	"os"

	"pearls-grader-python/config"
	"pearls-grader-python/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
