package day5

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(challenge *challenge.Input) int {
	seats := parseSeats(challenge)
	return seats[len(seats)-1]
}
