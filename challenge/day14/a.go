package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 14, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	var andMask int64 = 0
	var orMask int64 = 0

	mem := map[int]int64{}

	for op := range challenge.Lines() {
		parts := strings.Split(op, " = ")

		if parts[0] == "mask" {
			andMask, _ = strconv.ParseInt(strings.ReplaceAll(parts[1], "X", "1"), 2, 0)
			orMask, _ = strconv.ParseInt(strings.ReplaceAll(parts[1], "X", "0"), 2, 0)
		} else {
			addr := util.MustAtoI(strings.TrimSuffix(strings.TrimPrefix(parts[0], "mem["), "]"))
			value := int64(util.MustAtoI(parts[1]))

			value &= andMask
			value |= orMask

			mem[addr] = value
		}
	}

	answer := 0
	for _, v := range mem {
		answer += int(v)
	}

	return answer
}
