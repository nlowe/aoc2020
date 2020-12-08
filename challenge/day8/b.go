package day8

import (
	"fmt"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	instructions := newCpuFor(challenge)

	for i, instr := range instructions {
		orig := instr.op
		if instr.op == instrJmp {
			instructions[i].op = instrNop
		} else if instr.op == instrNop {
			instructions[i].op = instrJmp
		}

		if ok, acc := halts(instructions); ok {
			return acc
		}

		instructions[i].op = orig
	}

	panic("No solution")
}
