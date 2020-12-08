package day8

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
)

const (
	instrAcc = "acc"
	instrJmp = "jmp"
	instrNop = "nop"
)

type instruction struct {
	op  string
	rel int
}

func newCpuFor(challenge *challenge.Input) []instruction {
	lines := challenge.LineSlice()
	instructions := make([]instruction, len(lines))

	for i, instr := range lines {
		parts := strings.Split(instr, " ")

		instructions[i] = instruction{
			op:  parts[0],
			rel: util.MustAtoI(parts[1]),
		}
	}

	return instructions
}

func halts(instructions []instruction) (bool, int) {
	ip := 0
	acc := 0

	seen := map[int]struct{}{}

	for {
		instr := instructions[ip]
		if _, ok := seen[ip]; ok {
			return false, acc
		}

		seen[ip] = struct{}{}

		switch instr.op {
		case instrAcc:
			acc += instr.rel
			ip++
		case instrNop:
			ip++
		case instrJmp:
			ip += instr.rel
		default:
			panic(fmt.Errorf("[%d] unknown instruction: %s %d", ip, instr.op, instr.rel))
		}

		if ip == len(instructions) {
			return true, acc
		}
	}
}
