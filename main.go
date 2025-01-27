package main

import (
	"github.com/SchulteMarkus/sse-belmngr-hermine/cli"
	"os"
)

func main() {
	if err := cli.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
