package cmd

import (
	"fmt"
	"os"
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
	Use:   "del",
	Short: "Deletes the selected branches",
	Long:  `After selecting branches (based on their location - local or origin), will delete them`,
	Run: func(cmd *cobra.Command, args []string) {
		filter := util.GetConfigOrExit("listfilter")
		branches := listBranches(filter)

		// Commands to delete branches in local and remote
		local := "git branch -D "
		remote := "git push --delete origin "

		// Prompt which ones to delete
		filteredBranches := []*LocatedRef{}
		filteredBranchPrompts := []string{}
		for _, branch := range branches {
			if branch.Name().Short() == "master" {
				// We don't mess with master
				continue
			}

			filteredBranchPrompts = append(
				filteredBranchPrompts,
				fmt.Sprintf("%s\t%s", branch.LocationName(), branch.Name().Short()),
			)
			filteredBranches = append(filteredBranches, branch)
		}
		branchIdxs := prompt.SelectManyIndex("Select branches to delete", filteredBranchPrompts)

		// If nothing, we're done
		if len(branchIdxs) == 0 {
			fmt.Println("Nothing to delete")
			os.Exit(0)
		}

		// Cherry pick refs we're going to delete
		localList := []string{}
		remoteList := []string{}
		selectedRefs := []*LocatedRef{}
		for _, idx := range branchIdxs {
			ref := filteredBranches[idx]
			refName := ref.Name().Short()
			if ref.IsLocal() {
				localList = append(localList, refName)
			} else {
				remoteList = append(remoteList, refName)
			}
			selectedRefs = append(selectedRefs, ref)
		}

		// Warn user & confirm
		fmt.Println("About to delete:")
		for _, ref := range selectedRefs {
			fmt.Println(fmt.Sprintf("   %s\t%s", ref.LocationName(), ref.Name().Short()))
		}
		result := prompt.Bool("Are you sure")
		if !result {
			fmt.Println("Cancelled.")
			os.Exit(0)
		}

		// Go forth.
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
