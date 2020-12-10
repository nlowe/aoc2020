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
	// TODO: Figure out why removing the starting 0 works for
	//       the small test case but not the larger one
	adapters := append([]int{0}, parseAdapters(challenge)...)
	return permute(adapters, map[int]int{})
}

func permute(remaining []int, memo map[int]int) int {
	if len(remaining) == 1 {
		return 1
	}

	if count, ok := memo[remaining[0]]; ok {
		return count
	}

	count := 0
	for i := 1; i < len(remaining); i++ {
		if remaining[i]-remaining[0] <= 3 {
			count += permute(remaining[i:], memo)
		}
	}

	memo[remaining[0]] = count
	return count
}
