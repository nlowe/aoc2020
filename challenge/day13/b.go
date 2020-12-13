package day13

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/deanveloper/modmath/v1/bigmod"
	"github.com/nlowe/aoc2020/challenge"
	"github.com/nlowe/aoc2020/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 13, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %s\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) string {
	lines := challenge.Lines()
	<-lines

	var busses []bigmod.CrtEntry //nolint:prealloc
	rawSchedule := strings.Split(<-lines, ",")

	for off, rawId := range rawSchedule {
		if rawId == "x" {
			continue
		}

		id := util.MustAtoI(rawId)
		busses = append(busses, bigmod.CrtEntry{A: big.NewInt(int64(id - off)), N: big.NewInt(int64(id))})
	}

	return bigmod.SolveCrtMany(busses).String()
}
