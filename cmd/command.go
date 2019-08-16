package cmd

import "github.com/spf13/cobra"

// Init all the commands; consider this the entry point
func Init(root *cobra.Command) {
	commands := []*cobra.Command{
		newStartCmd(),
		newDelCmd(),
		newGenerateCmd(),
		newListCmd(),
	}

	for _, command := range commands {
		root.AddCommand(command)
	}
}
