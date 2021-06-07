package main

import (
	"os"

	"github.com/mconcat/ibc-eth/cmd/ibc-ethd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
