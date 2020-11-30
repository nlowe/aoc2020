package main

import (
	"os"

	"github.com/nlowe/aoc2020/challenge/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
