package day10

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 10, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	return permute(0, parseAdapters(challenge), map[int]int{})
}

func permute(previous int, remaining []int, memo map[int]int) int {
	if len(remaining) == 1 {
		return 1
	}

	if count, ok := memo[previous]; ok {
		return count
	}

	// Assume remaining[0] is required
	count := permute(remaining[0], remaining[1:], memo)

	// If it's optional, also add the permutations without it
	if remaining[1]-previous <= 3 {
		count += permute(previous, remaining[1:], memo)
	}

	memo[previous] = count
	return count
}
