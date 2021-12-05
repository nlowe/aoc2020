package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

const maxDepth = 10

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 19, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	rules := map[int]rule{}

	buildingRules := true
	matches := 0

	var regex *regexp.Regexp

	for line := range challenge.Lines() {
		if line == "" {
			buildingRules = false

			rules[8] = rule(strings.Join(hack(func(i int) string {
				return strings.Repeat("42 ", i)
			}), "|"))
			rules[11] = rule(strings.Join(hack(func(i int) string {
				return strings.Repeat("42 ", i) + " " + strings.Repeat("31 ", i)
			}), "|"))

			regex = regexp.MustCompile("^" + rules[0].toRegex(rules) + "$")

			continue
		}

		if buildingRules {
			parts := strings.Split(line, ": ")
			rules[util.MustAtoI(parts[0])] = rule(parts[1])
		} else if line != "" && regex != nil {
			if regex.Match([]byte(line)) {
				matches++
			}
		}
	}

	return matches
}

func hack(gen func(int) string) []string {
	result := make([]string, maxDepth-1)

	for i := 1; i < maxDepth; i++ {
		result[i-1] = gen(i)
	}

	return result
}
