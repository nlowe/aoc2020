package day1

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	entries := parsePasswords(challenge)

	for i := range entries {
		for j := i + 1; j < len(entries)-1; j++ {
			for k := j + 1; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == 2020 {
					return entries[i] * entries[j] * entries[k]
				}
			}
		}
	}

	panic("no solution")
}
