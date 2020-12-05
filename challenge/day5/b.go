package day5

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	seats := parseSeats(challenge)

	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] == 2 {
			return seats[i] + 1
		}
	}

	panic("no solution")
}
