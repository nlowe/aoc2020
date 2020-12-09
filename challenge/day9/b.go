package day9

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	return findWeakness(challenge, challengePreambleLength)
}

func findWeakness(challenge *challenge.Input, preamble int) int {
	numbers, off, n := findInvalid(challenge, preamble)

loop:
	for i := 0; i < off-1; i++ {
		smallest := numbers[i]
		largest := numbers[i]

		sum := numbers[i]
		for j := i + 1; j < off; j++ {
			sum += numbers[j]

			if numbers[j] < smallest {
				smallest = numbers[j]
			} else if numbers[j] > largest {
				largest = numbers[j]
			}

			if sum == n {
				return smallest + largest
			} else if sum > n {
				continue loop
			}
		}
	}

	panic("no solution")
}
