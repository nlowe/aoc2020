package day13

import (
	"fmt"
	"math"
	"strings"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 13, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	lines := challenge.Lines()
	earliest := util.MustAtoI(<-lines)

	var busses []int //nolint:prealloc
	schedule := strings.Split(<-lines, ",")

	for _, id := range schedule {
		if id == "x" {
			continue
		}

		busses = append(busses, util.MustAtoI(id))
	}

	bestBus := 0
	best := math.MaxInt64
	for _, bus := range busses {
		wait := bus - (earliest % bus)
		if wait < best {
			best = wait
			bestBus = bus
		}
	}

	return bestBus * best
}
