package day1

import (
	"fmt"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {
	var entries []int

	for v := range challenge.Lines() {
		entries = append(entries, util.MustAtoI(v))
	}

	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			if entries[i]+entries[j] == 2020 {
				return entries[i] * entries[j]
			}
		}
	}

	panic("no solution")
}
