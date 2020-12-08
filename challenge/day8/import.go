package day8

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "8",
		Short: "Problems for Day 8",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())

	root.AddCommand(day)
}
