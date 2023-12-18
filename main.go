package main

import (
	"os"

	"github.com/daoprover/listener-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
