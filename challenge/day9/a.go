package day9

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

const challengePreambleLength = 25

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	_, _, n := findInvalid(challenge, challengePreambleLength)
	return n
}

func findInvalid(challenge *challenge.Input, preamble int) ([]int, int, int) {
	lines := challenge.LineSlice()
	numbers := make([]int, len(lines))

	for i, line := range lines {
		n := util.MustAtoI(line)

		numbers[i] = n
	}

loop:
	for i := preamble; i < len(numbers); i++ {
		for j := i - preamble; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if numbers[j]+numbers[k] == numbers[i] {
					continue loop
				}
			}
		}

		return numbers, i, numbers[i]
	}

	panic("no solution")
}
