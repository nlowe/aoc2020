package day12

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

const (
	north = iota
	east
	south
	west
)

const (
	opNorth       = 'N'
	opEast        = 'E'
	opSouth       = 'S'
	opWest        = 'W'
	opRotateLeft  = 'L'
	opRotateRight = 'R'
	opForward     = 'F'
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 12, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	x := 0
	y := 0
	face := east

	for instruction := range challenge.Lines() {
		op := instruction[0]
		distance := util.MustAtoI(instruction[1:])

		if op == opForward {
			switch face {
			case north:
				op = opNorth
			case east:
				op = opEast
			case south:
				op = opSouth
			case west:
				op = opWest
			default:
				panic(fmt.Errorf("unknown face: %d", face))
			}
		}

		switch op {
		case opNorth:
			y += distance
		case opSouth:
			y -= distance
		case opEast:
			x += distance
		case opWest:
			x -= distance
		case opRotateLeft:
			face = util.IntSafeMod(face-distance/90, 4)
		case opRotateRight:
			face = util.IntSafeMod(face+distance/90, 4)
		default:
			panic(fmt.Errorf("unknown op: %s", string(op)))
		}
	}

	return util.ManhattanDistance(0, 0, x, y)
}
