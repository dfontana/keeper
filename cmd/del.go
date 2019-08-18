package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dfontana/keeper/prompt"
	"github.com/dfontana/keeper/util"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
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
		r := util.OpenRepoOrExit()
		branches := listBranches(r, filter)

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
		localRefs := []*LocatedRef{}
		remoteRefs := []*LocatedRef{}
		for _, idx := range branchIdxs {
			ref := filteredBranches[idx]
			if ref.IsLocal() {
				localRefs = append(localRefs, ref)
			} else {
				remoteRefs = append(remoteRefs, ref)
			}
		}

		// Warn user & confirm
		fmt.Println("About to delete:")
		for _, ref := range remoteRefs {
			fmt.Println(fmt.Sprintf("   %s\t%s", ref.LocationName(), ref.Name().Short()))
		}
		for _, ref := range localRefs {
			fmt.Println(fmt.Sprintf("   %s\t%s", ref.LocationName(), ref.Name().Short()))
		}
		result := prompt.Bool("Are you sure")
		if !result {
			fmt.Println("Cancelled.")
			os.Exit(0)
		}

		// Go forth.
		deleteRemote(r, remoteRefs)
		deleteLocal(r, localRefs)
	},
}

func deleteRemote(r *git.Repository, refs []*LocatedRef) {
	// TODO figure out how to get SSH auth working under git.Push
	// updates := []config.RefSpec{}
	// for _, ref := range refs {
	// 	updates = append(updates, config.RefSpec(":refs/heads/"+ref.Name().Short()))
	// }
	// auth, err := ssh.NewSSHAgentAuth("")
	// util.CheckSafeExit("Failed to get ssh auth", err)
	// err = r.Push(&git.PushOptions{
	// 	Auth:     auth,
	// 	RefSpecs: updates,
	// 	Progress: os.Stdout,
	// })
	// util.CheckSafeExit("Failed to clear remotes, stopping", err)
	names := []string{}
	for _, ref := range refs {
		names = append(names, ref.Name().Short())
	}
	cmd := fmt.Sprintf("git push --delete origin %s", strings.Join(names, " "))
	err := util.RunString(cmd)
	util.CheckSafeExit("Can't complete remote", err)
}

func deleteLocal(r *git.Repository, refs []*LocatedRef) {
	for _, ref := range refs {
		err := r.Storer.RemoveReference(ref.Name())
		util.CheckSafeExit("Failed to clear local branch, stopping", err)
	}
}

func newDelCmd() *cobra.Command {
	return delCmd
}
