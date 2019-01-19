package cmd

import (
	"fmt"
	"strings"

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

		// Commands to validate if branches exist locally and remote
		testLocal := "git show-ref --verify --quiet refs/heads/"
		testRemote := "git ls-remote --heads --exit-code origin"

		// Commands to delete branches in local and remote
		local := "git branch -D"
		remote := "git push --delete origin"

		// Cofirm each branch with the user, to be sure
		// We'll only consider items that exist in either location irrespective
		// to the other
		localList := []string{}
		remoteList := []string{}
		for _, branch := range args {
			ans := PromptBool(fmt.Sprintf("Delete %s", branch))
			if ans {
				if err := RunString(testLocal + branch); err == nil {
					localList = append(localList, branch)
				}
				if err := RunString(testRemote + branch); err == nil {
					remoteList = append(remoteList, branch)
				}
			} else {
				fmt.Printf("Skipping %s\n", branch)
			}
		}

		if err := RunString(local + strings.Join(localList, " ")); err != nil {
			return
		}

		if err := RunString(remote + strings.Join(remoteList, " ")); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
