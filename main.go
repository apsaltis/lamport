package main

import (
	"os"

	"github.com/Distributed-Computing-Denver/lamport/cli"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
