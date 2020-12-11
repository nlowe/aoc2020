package day11

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 11, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	layout := challenge.TileMap()

	for step(layout, countNeighbors, 4) {
	}

	return countOccupied(layout)
}

func countNeighbors(seats *challenge.TileMap, x, y int) int {
	return countOccupiedNeighbors(seats, x, y, func(seats *challenge.TileMap, x, y, dx, dy int) (rune, bool) {
		return seats.TileAt(x+dx, y+dy)
	})
}
