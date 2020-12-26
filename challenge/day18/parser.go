package day18

import (
	"fmt"
	"strconv"
	"strings"
)

var normalizer = strings.NewReplacer("(", " ( ", ")", " ) ")

func evaluate(eqn string, precedenceMap map[string]int) int {
	// Use https://www.rosettacode.org/wiki/Parsing/Shunting-yard_algorithm#Go
	// to convert to RPN, then evaluate it
	return evaluateRPN(toRPN(strings.Fields(normalizer.Replace(eqn)), precedenceMap))
}

// toRPN converts the tokens of an infix equation to postfix notation
func toRPN(tokens []string, precedenceMap map[string]int) (result []string) {
	var stack []string

	for _, token := range tokens {
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			var op string
			for len(stack) > 0 {
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // Discard "("
				}

				result = append(result, op)
			}
		default:
			if op1precedence, ok := precedenceMap[token]; ok { // Operator
				for len(stack) > 0 {
					top := stack[len(stack)-1]
					if op2precedence, ok := precedenceMap[top]; !ok || op1precedence > op2precedence {
						break
					}

					stack, result = stack[:len(stack)-1], append(result, top)
				}

				stack = append(stack, token)
			} else { // Operand
				result = append(result, token)
			}
		}
	}

	// drain remaining stack
	for len(stack) > 0 {
		result, stack = append(result, stack[len(stack)-1]), stack[:len(stack)-1]
	}

	return
}

// evaluateRPN takes an equation in postfix notation and returns the evaluation of it
func evaluateRPN(tokens []string) int {
	var values []int

	for _, t := range tokens {
		v, ok := strconv.Atoi(t)

		if ok == nil {
			values = append(values, v)
		} else {
			if values == nil || len(values) < 2 {
				panic(fmt.Errorf("not enough values remaining for operator %s: %+v", t, values))
			}

			var v1, v2 int
			v1, values = values[len(values)-1], values[:len(values)-1]
			v2, values = values[len(values)-1], values[:len(values)-1]

			switch t {
			case "*":
				values = append(values, v1*v2)
			case "+":
				values = append(values, v1+v2)
			default:
				panic(fmt.Errorf("unsupported operator or unknown token: %s", t))
			}
		}
	}

	if values == nil || len(values) != 1 {
		panic(fmt.Errorf("expected exactly 1 value remaining but had %+v", values))
	}

	return values[0]
}
