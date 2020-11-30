package example

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "example",
		Short: "Problems for Day Example",
	}

	day.AddCommand(aCommand())

	root.AddCommand(day)
}
