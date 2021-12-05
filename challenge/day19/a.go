package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlowe/aoc2020/util"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 19, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type rule string

func (r rule) toRegex(rules map[int]rule) string {
	switch {
	case r == `"a"`:
		return "a"
	case r == `"b"`:
		return "b"
	case strings.Contains(string(r), "|"):
		result := strings.Builder{}
		for _, nested := range strings.Split(string(r), "|") {
			result.WriteString("(?:")
			result.WriteString(rule(nested).toRegex(rules))
			result.WriteString(")|")
		}

		return strings.TrimRight(result.String(), "|")
	case strings.Contains(string(r), " "):
		result := strings.Builder{}
		for _, seq := range strings.Fields(string(r)) {
			result.WriteString("(?:")
			result.WriteString(rule(seq).toRegex(rules))
			result.WriteString(")")
		}

		return result.String()
	default:
		return rules[util.MustAtoI(string(r))].toRegex(rules)
	}
}

func partA(challenge *challenge.Input) int {
	rules := map[int]rule{}

	buildingRules := true
	matches := 0

	var regex *regexp.Regexp

	for line := range challenge.Lines() {
		if line == "" {
			buildingRules = false

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
