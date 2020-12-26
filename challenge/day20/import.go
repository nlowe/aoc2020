package day20

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "20",
		Short: "Problems for Day 20",
	}

	day.AddCommand(aCommand())
	day.AddCommand(bCommand())

	root.AddCommand(day)
}
