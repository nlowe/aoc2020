package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

type opFn func(left, right int) int

var (
	opAdd = func(left, right int) int { return left + right }
	opMul = func(left, right int) int { return left * right }
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 18, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	normalizer := strings.NewReplacer("(", " ( ", ")", " ) ")

	sum := 0
	for eqn := range challenge.Lines() {
		sum += parseTerm(strings.Fields(normalizer.Replace(eqn)))
	}

	return sum
}

func parseTerm(tokens []string) int {
	left := 0
	var op opFn

	skip := 0
	for off, t := range tokens {
		if skip > 0 {
			skip--
			continue
		}

		v, err := strconv.Atoi(t)

		switch {
		case t == "(":
			skip = closingParenthesis(tokens[off:])

			v = parseTerm(tokens[off+1 : off+skip])
			if op == nil {
				left = v
			} else {
				left = op(left, v)
				op = nil
			}
		case t == "*":
			op = opMul
		case t == "+":
			op = opAdd
		case err == nil:
			if op == nil {
				left = v
			} else {
				left = op(left, v)
				op = nil
			}
		default:
			panic(fmt.Errorf("unknown token: %s", t))
		}
	}

	return left
}

func closingParenthesis(tokens []string) int {
	if len(tokens) == 0 || tokens[0] != "(" {
		panic(fmt.Errorf("no matching parenthesis remaining in %+v", tokens))
	}

	count := 0
	for i, t := range tokens {
		if t == "(" {
			count++
		} else if t == ")" {
			count--

			if count == 0 {
				return i
			}
		}
	}

	panic(fmt.Errorf("no matching parenthesis remaining in %+v", tokens))
}
