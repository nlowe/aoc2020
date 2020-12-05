package day1

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func parsePasswords(challenge *challenge.Input) []int {
	lines := challenge.LineSlice()
	entries := make([]int, len(lines))

	for i, v := range lines {
		entries[i] = util.MustAtoI(v)
	}

	return entries
}

func partA(challenge *challenge.Input) int {
	entries := parsePasswords(challenge)

	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			if entries[i]+entries[j] == 2020 {
				return entries[i] * entries[j]
			}
		}
	}

	panic("no solution")
}
