package day23

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "23",
		Short: "Problems for Day 23",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())

	root.AddCommand(day)
}
