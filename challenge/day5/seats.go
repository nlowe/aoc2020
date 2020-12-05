package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2020/challenge"
)

const (
	tokenFront = "F"
	tokenBack  = "B"
	tokenLeft  = "L"
	tokenRight = "R"
)

var replacer = strings.NewReplacer(tokenFront, "0", tokenBack, "1", tokenLeft, "0", tokenRight, "1")

func seatID(line string) int {
	id, err := strconv.ParseInt(replacer.Replace(line), 2, 0)
	if err != nil {
		panic(fmt.Errorf("failed to parse seat %s: %w", line, err))
	}

	return int(id)
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
