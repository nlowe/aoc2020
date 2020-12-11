package day11

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 11, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	layout := challenge.TileMap()

	for step(layout, countVisibleNeighbors, 5) {
	}

	return countOccupied(layout)
}

func countVisibleNeighbors(seats *challenge.TileMap, x, y int) int {
	return countOccupiedNeighbors(seats, x, y, raycast)
}

func raycast(seats *challenge.TileMap, sx, sy, dx, dy int) (rune, bool) {
	x := sx + dx
	y := sy + dy
	for {
		t, ok := seats.TileAt(x, y)
		if !ok {
			return ' ', false
		}

		if t != tileFloor {
			return t, true
		}

		x += dx
		y += dy
	}
}
