package day16

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 16, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	rules, _, nearbyTickets := parseTickets(challenge)

	errorRate := 0
	for _, t := range nearbyTickets {
	check:
		for _, raw := range t {
			for _, r := range rules {
				if r.satisfies(raw) {
					continue check
				}
			}

			errorRate += raw
		}
	}

	return errorRate
}
