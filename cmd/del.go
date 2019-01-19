package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del <branch_name> <branch_name> <...>",
	Short: "Deletes a branch both locally and remotely",
	Long:  `Given a list of branch names (without the origin/), will delete each one both locally and remotely.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		// Cofirm each branch with the user, to be sure
		confirmedList := []string{}
		for _, branch := range args {
			ans := PromptBool(fmt.Sprintf("Delete %s", branch))
			if ans {
				confirmedList = append(confirmedList, branch)
			} else {
				fmt.Printf("Skipping %s\n", branch)
			}
		}

		local := []string{
			"git",
			"branch",
			"-D",
		}
		remote := []string{
			"git",
			"push",
			"--delete",
			"origin",
		}

		if err := Run(append(local, confirmedList...)); err != nil {
			return
		}

		if err := Run(append(remote, confirmedList...)); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
