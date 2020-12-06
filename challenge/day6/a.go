package day6

import (
	"fmt"
	"math/bits"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	answer := 0

	var wip uint = 0
	for line := range challenge.Lines() {
		if line == "" {
			answer += bits.OnesCount(wip)
			wip = 0
			continue
		}

		for _, response := range line {
			wip |= 1 << (response - 'a')
		}
	}

	return answer + bits.OnesCount(wip)
}
