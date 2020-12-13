package day13

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "13",
		Short: "Problems for Day 13",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())

	root.AddCommand(day)
}
