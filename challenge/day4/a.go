package day4

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	lines := challenge.Lines()

	valid := 0
	for {
		p := parsePassport(lines)

		if p == (passport{}) {
			break
		}

		if p.valid() {
			valid++
		}
	}

	return valid
}
