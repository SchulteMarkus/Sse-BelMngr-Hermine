package main

import (
	"os"
	"schulte.dev/sse-belmngr-hermine/cli"
)

func main() {
	if err := cli.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
