package day16

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 16, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	rules, myTicket, nearbyTickets := parseTickets(challenge)

	var validTickets []ticket

search:
	for _, t := range nearbyTickets {
	check:
		for _, raw := range t {
			for _, r := range rules {
				if r.satisfies(raw) {
					continue check
				}
			}

			// invalid
			continue search
		}

		validTickets = append(validTickets, t)
	}

	if len(validTickets) == 0 {
		panic("no valid tickets")
	}

	// For each rule, which columns could it be?
	fieldLocationCandidates := map[string]map[int]struct{}{}
	for _, r := range rules {
		if _, ok := fieldLocationCandidates[r.name]; !ok {
			fieldLocationCandidates[r.name] = map[int]struct{}{}
		}

	fieldNumberSearch:
		for fieldNumber := 0; fieldNumber < len(myTicket); fieldNumber++ {
			for _, t := range validTickets {
				if !r.satisfies(t[fieldNumber]) {
					continue fieldNumberSearch
				}
			}

			fieldLocationCandidates[r.name][fieldNumber] = struct{}{}
		}
	}

	// while there are more candidates
	fieldLocations := map[string]int{}
	for len(fieldLocationCandidates) > 0 {
		for idx, candidates := range fieldLocationCandidates {
			//   find a candidate with only one valid location
			if len(candidates) == 1 {
				//   remember its location
				location := -1
				for location = range candidates {
				}

				fieldLocations[idx] = location

				//   remove its location from consideration of all other candidates
				for k := range fieldLocationCandidates {
					delete(fieldLocationCandidates[k], location)
				}

				delete(fieldLocationCandidates, idx)
				break
			}
		}
	}

	checksum := 1
	for name, location := range fieldLocations {
		if strings.HasPrefix(name, "departure") {
			checksum *= myTicket[location]
		}
	}

	return checksum
}
