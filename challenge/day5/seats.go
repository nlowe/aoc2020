package day5

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2020/challenge"
)

func seatID(line string) int {
	minRow := 0
	maxRow := rows

	minColumn := 0
	maxColumn := columns

	for _, token := range line {
		rowDelta := (maxRow-minRow)/2 + 1
		columnDelta := (maxColumn-minColumn)/2 + 1

		switch token {
		case tokenFront:
			maxRow -= rowDelta
		case tokenBack:
			minRow += rowDelta
		case tokenLeft:
			maxColumn -= columnDelta
		case tokenRight:
			minColumn += columnDelta
		default:
			panic(fmt.Errorf("unknown token parsing assignment %s: %s", line, string(token)))
		}
	}

	return maxRow*(columns+1) + maxColumn
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
