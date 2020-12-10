package day10

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 10, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	adapters := parseAdapters(challenge)

	d1 := 0
	d3 := 1 // Adapter in bag is automatically highest+3

	previous := 0
	for _, adapter := range adapters {
		delta := adapter - previous
		if delta == 3 {
			d3++
		} else if delta == 1 {
			d1++
		}

		previous = adapter
	}

	return d1 * d3
}

func parseAdapters(challenge *challenge.Input) []int {
	lines := challenge.LineSlice()
	adapters := make([]int, len(lines))
	for i, line := range lines {
		adapters[i] = util.MustAtoI(line)
	}

	sort.Ints(adapters)
	return adapters
}
