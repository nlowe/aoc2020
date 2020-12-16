package day16

import (
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
)

const (
	parseStateRules = iota
	parseStateMyTicket
	parseStateNearbyTickets
)

type ticket []int

type validRange struct {
	low  int
	high int
}

func (v validRange) satisfies(i int) bool {
	return i >= v.low && i <= v.high
}

type rule struct {
	name string

	a validRange
	b validRange
}

func (r rule) satisfies(i int) bool {
	return r.a.satisfies(i) || r.b.satisfies(i)
}

func parseTickets(challenge *challenge.Input) ([]rule, ticket, []ticket) {
	var rules []rule
	var myTicket ticket
	var nearbyTickets []ticket

	state := parseStateRules
	for line := range challenge.Lines() {
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "your ticket"):
			state = parseStateMyTicket
			continue
		case strings.HasPrefix(line, "nearby tickets"):
			state = parseStateNearbyTickets
			continue
		}

		switch state {
		case parseStateRules:
			rules = append(rules, parseRule(line))
		case parseStateMyTicket:
			myTicket = parseTicket(line)
		case parseStateNearbyTickets:
			nearbyTickets = append(nearbyTickets, parseTicket(line))
		}
	}

	return rules, myTicket, nearbyTickets
}

func parseRule(line string) rule {
	parts := strings.Split(line, ": ")
	result := rule{name: parts[0]}

	ranges := strings.Split(parts[1], " or ")
	a := strings.Split(ranges[0], "-")
	b := strings.Split(ranges[1], "-")

	result.a = validRange{low: util.MustAtoI(a[0]), high: util.MustAtoI(a[1])}
	result.b = validRange{low: util.MustAtoI(b[0]), high: util.MustAtoI(b[1])}

	return result
}

func parseTicket(line string) ticket {
	parsed := strings.Split(line, ",")
	var result ticket = make([]int, len(parsed))

	for i, v := range parsed {
		result[i] = util.MustAtoI(v)
	}

	return result
}
