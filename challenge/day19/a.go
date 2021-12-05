package day19

import (
	"fmt"
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

type Rule interface {
	// Match returns the remaining portion of the string and whether or not the rule matched
	Match(string) (string, bool)
	// Link un-refs rules once all rules are known
	Link(map[int]Rule)
}

// literal is a Rule that matches a single rune
type literal rune

func (l literal) Match(s string) (string, bool) {
	if len(s) >= 1 && rune(s[0]) == rune(l) {
		return s[1:], true
	}

	return s, false
}
func (l literal) Link(map[int]Rule) {}

// or is a Rule that matches when either of its sub rules are matched
type or struct {
	left  Rule
	right Rule
}

func (o or) Match(s string) (string, bool) {
	if remaining, match := o.left.Match(s); match {
		return remaining, true
	} else if remaining, match := o.right.Match(s); match {
		return remaining, true
	}

	return s, false
}
func (o or) Link(rules map[int]Rule) {
	o.left.Link(rules)
	o.right.Link(rules)
}

// sequence is a Rule that matches if all rules in the sequence match in order
type sequence []Rule

func (s sequence) Match(str string) (string, bool) {
	remaining := str
	match := false
	for _, rule := range s {
		remaining, match = rule.Match(remaining)
		if !match {
			return str, false
		}
	}

	return remaining, true
}
func (s sequence) Link(rules map[int]Rule) {
	for _, r := range s {
		r.Link(rules)
	}
}

// ref is a Rule that references another rule in the rule set
type ref struct {
	n   int
	ref Rule
}

func (r ref) Match(s string) (string, bool) {
	if r.ref == nil {
		panic(fmt.Errorf("reference rule %d was not linked", r.n))
	}

	return r.ref.Match(s)
}
func (r *ref) Link(rules map[int]Rule) {
	r.ref = rules[r.n]
}

func partA(challenge *challenge.Input) int {
	rules := map[int]Rule{}

	buildingRules := true
	matches := 0

	for line := range challenge.Lines() {
		if line == "" {
			buildingRules = false

			// Link all ref rules
			for _, r := range rules {
				r.Link(rules)
			}

			continue
		}

		if buildingRules {
			n, r := parse(line)
			rules[n] = r
		} else {
			if left, match := rules[0].Match(line); match && left == "" {
				matches++
			}
		}
	}

	return matches
}

func parse(line string) (int, Rule) {
	parts := strings.Split(line, ":")

	n := util.MustAtoI(parts[0])

	rawRule := strings.TrimSpace(parts[1])

	if strings.Contains(rawRule, "|") {
		parts = strings.Split(rawRule, "|")

		return n, or{
			left:  parseSequence(parts[0]),
			right: parseSequence(parts[1]),
		}
	} else if rawRule == `"a"` {
		return n, literal('a')
	} else if rawRule == `"b"` {
		return n, literal('b')
	} else {
		// It's a sequence
		return n, parseSequence(rawRule)
	}
}

func parseSequence(rawRule string) sequence {
	rules := strings.Fields(rawRule)

	var seq sequence
	for _, rule := range rules {
		seq = append(seq, &ref{n: util.MustAtoI(rule)})
	}

	return seq
}
