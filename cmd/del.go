package cmd

import (
	"fmt"
	"strings"

	"github.com/dfontana/keeper/prompt"
	"github.com/dfontana/keeper/util"
	"github.com/spf13/cobra"
)

func mapAry(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// DelCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del <branch_name> <branch_name> <...>",
	Short: "Deletes a branch both locally and remotely",
	Long:  `Given a list of branch names (without the origin/), will delete each one both locally and remotely.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		// Commands to validate if branches exist locally and remote
		testLocal := "git show-ref --verify --quiet refs/heads/"
		testRemote := "git ls-remote --heads --exit-code origin "

		// Commands to delete branches in local and remote
		local := "git branch -D "
		remote := "git push --delete origin "

		// Cofirm each branch with the user, to be sure
		// We'll only consider items that exist in either location irrespective
		// to the other
		localList := []string{}
		remoteList := []string{}
		for _, branch := range args {
			ans := prompt.Bool(fmt.Sprintf("Delete %s", branch))
			if ans {
				if err := util.RunString(testLocal + branch); err == nil {
					localList = append(localList, branch)
				}
				if err := util.RunString(testRemote + branch); err == nil {
					remoteList = append(remoteList, branch)
				}
			} else {
				fmt.Printf("Skipping %s\n", branch)
			}
		}
		fmt.Println("Deleting:")
		fmt.Println(strings.Join(mapAry(localList, func(v string) string {
			return "\t[Local]: " + v
		}), " "))
		fmt.Println(strings.Join(mapAry(remoteList, func(v string) string {
			return "\t[Remote]: " + v
		}), " "))

		if len(localList) > 0 {
			if err := util.RunString(local + strings.Join(localList, " ")); err != nil {
				return
			}
		}

		if len(remoteList) > 0 {
			if err := util.RunString(remote + strings.Join(remoteList, " ")); err != nil {
				return
			}
		}
	},
}

func newDelCmd() *cobra.Command {
	return delCmd
}
