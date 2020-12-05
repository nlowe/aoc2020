package day4

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	lines := challenge.Lines()

	valid := 0
	for {
		p := parsePassport(lines)

		if p == (passport{}) {
			break
		}

		if p.strictlyValid() {
			valid++
		}
	}

	return valid
}
