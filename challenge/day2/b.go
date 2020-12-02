package day2

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func (p policy) strictlyValid(password string) bool {
	return (rune(password[p.min-1]) == p.target) != (rune(password[p.max-1]) == p.target)
}

func b(challenge *challenge.Input) int {
	valid := 0

	for line := range challenge.Lines() {
		parts := strings.Split(line, ": ")

		if parsePolicy(parts[0]).strictlyValid(parts[1]) {
			valid++
		}
	}

	return valid
}
