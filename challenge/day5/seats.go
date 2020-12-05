package day5

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2020/challenge"
)

const (
	tokenFront = 'F'
	tokenBack  = 'B'
	tokenLeft  = 'L'
	tokenRight = 'R'
)

func seatID(line string) int {
	id := 0

	for _, token := range line {
		id <<= 1
		switch token {
		case tokenFront:
		case tokenBack:
			id |= 1
		case tokenLeft:
		case tokenRight:
			id |= 1
		default:
			panic(fmt.Errorf("unknown token parsing assignment %s: %s", line, string(token)))
		}
	}

	return id
}

func parseSeats(challenge *challenge.Input) []int {
	lines := challenge.LineSlice()
	seatIDs := make([]int, len(lines))

	for i, seat := range lines {
		seatIDs[i] = seatID(seat)
	}

	sort.Ints(seatIDs)
	return seatIDs
}
