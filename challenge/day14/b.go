package day14

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 14, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	mask := ""
	mem := map[string]int{}

	for op := range challenge.Lines() {
		parts := strings.Split(op, " = ")

		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			addr := util.MustAtoI(strings.TrimSuffix(strings.TrimPrefix(parts[0], "mem["), "]"))
			value := util.MustAtoI(parts[1])

			for _, generatedAddr := range permuteMask("", mask, fmt.Sprintf("%036b", addr)) {
				mem[generatedAddr] = value
			}
		}
	}

	answer := 0
	for _, v := range mem {
		answer += v
	}

	return answer
}

func permuteMask(mask, remaining, addr string) []string {
	if len(remaining) == 0 {
		return []string{mask}
	}

	switch remaining[0] {
	case '0':
		return permuteMask(mask+string(addr[len(mask)]), remaining[1:], addr)
	case '1':
		return permuteMask(mask+"1", remaining[1:], addr)
	case 'X':
		return append(permuteMask(mask+"0", remaining[1:], addr), permuteMask(mask+"1", remaining[1:], addr)...)
	default:
		panic(fmt.Errorf("unknown bitmask type %s", string(remaining[0])))
	}
}
