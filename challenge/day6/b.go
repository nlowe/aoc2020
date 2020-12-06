package day6

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	answer := 0

	var wip uint = math.MaxUint64
	for line := range challenge.Lines() {
		if line == "" {
			answer += bits.OnesCount(wip)
			wip = math.MaxUint64
			continue
		}

		var mask uint = 0
		for _, response := range line {
			mask |= 1 << (response - 'a')
		}

		wip &= mask
	}

	return answer + bits.OnesCount(wip)
}
