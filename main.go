package main

import (
	"os"

	"github.com/rabbits0209/komari-agent/cmd"
)

func main() {
	cmd.Execute()
	os.Exit(0)
}
