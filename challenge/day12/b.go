package day12

import (
	"fmt"
	"math"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 12, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	x := 0
	y := 0

	wx := 10
	wy := 1

	for instruction := range challenge.Lines() {
		op := instruction[0]
		distance := util.MustAtoI(instruction[1:])

		switch op {
		case opNorth:
			wy += distance
		case opSouth:
			wy -= distance
		case opEast:
			wx += distance
		case opWest:
			wx -= distance
		case opRotateLeft:
			t := rad(distance)
			c, s := math.Cos(t), math.Sin(t)
			wx, wy = wx*int(c)-wy*int(s), wx*int(s)+wy*int(c)
		case opRotateRight:
			t := rad(distance)
			c, s := math.Cos(t), math.Sin(t)
			wx, wy = wx*int(c)+wy*int(s), wx*(-1*int(s))+wy*int(c)
		case 'F':
			x += wx * distance
			y += wy * distance
		default:
			panic(fmt.Errorf("unknown op: %s", string(op)))
		}
	}

	return util.ManhattanDistance(0, 0, x, y)
}

func rad(d int) float64 {
	return float64(d) * (math.Pi / 180.0)
}
