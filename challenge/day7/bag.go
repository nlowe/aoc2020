package day7

import (
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
)

var myBag = bagKind{
	modifier: "shiny",
	color:    "gold",
}

type bagKind struct {
	modifier string
	color    string
}

type baggageRules map[bagKind]map[bagKind]int

func parseBagRules(challenge *challenge.Input) baggageRules {
	results := baggageRules{}

	for rawRule := range challenge.Lines() {
		parts := strings.Split(rawRule, " bags contain ")
		idParts := strings.Split(parts[0], " ")

		id := bagKind{
			modifier: strings.TrimSpace(idParts[0]),
			color:    strings.TrimSpace(idParts[1]),
		}

		results[id] = map[bagKind]int{}

		if parts[1] != "no other bags." {
			for _, inner := range strings.Split(parts[1], ", ") {
				parts := strings.Split(inner, " ")
				results[id][bagKind{modifier: parts[1], color: parts[2]}] = util.MustAtoI(parts[0])
			}
		}
	}

	return results
}

func (r baggageRules) canContain(start, target bagKind) bool {
	if n, ok := r[start][target]; ok && n > 0 {
		return true
	}

	for k := range r[start] {
		if r.canContain(k, target) {
			return true
		}
	}

	return false
}

func (r baggageRules) cost(k bagKind) int {
	totalCost := 1

	for rule, n := range r[k] {
		totalCost += n * r.cost(rule)
	}

	return totalCost
}
