package day2

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type policy struct {
	target rune
	min    int
	max    int
}

func parsePolicy(line string) policy {
	parts := strings.Split(line, " ")
	minMax := strings.Split(parts[0], "-")

	return policy{
		target: rune(parts[1][0]),
		min:    util.MustAtoI(minMax[0]),
		max:    util.MustAtoI(minMax[1]),
	}
}

func (p policy) valid(password string) bool {
	count := 0

	for _, r := range password {
		if r == p.target {
			count++
		}
	}

	return count >= p.min && count <= p.max
}

func partA(challenge *challenge.Input) int {
	valid := 0

	for line := range challenge.Lines() {
		parts := strings.Split(line, ": ")

		if parsePolicy(parts[0]).valid(parts[1]) {
			valid++
		}
	}

	return valid
}
