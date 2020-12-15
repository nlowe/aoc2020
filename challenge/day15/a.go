package day15

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 15, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return playTo(challenge, 2020)
}

func playTo(challenge *challenge.Input, turnCount int) int {
	spokenOnTurns := map[int][]int{}

	turn := 1
	lastSpoken := 0

	parts := strings.Split(<-challenge.Lines(), ",")
	for _, n := range parts {
		lastSpoken = util.MustAtoI(n)
		spokenOnTurns[lastSpoken] = []int{turn}
		turn++
	}

	for turn <= turnCount {
		turns := spokenOnTurns[lastSpoken]
		if len(turns) == 1 {
			lastSpoken = 0
		} else {
			lastSpoken = turns[len(turns)-1] - turns[len(turns)-2]
		}

		spokenOnTurns[lastSpoken] = append(spokenOnTurns[lastSpoken], turn)
		turn++
	}

	return lastSpoken
}
