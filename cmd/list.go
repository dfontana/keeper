package cmd

import (
	"fmt"
	"regexp"

	"github.com/dfontana/keeper/util"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

var insensitive bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <search string>",
	Short: "List branches based on the search string",
	Long:  `Can search over the author name or branch name. If no search string is given, then this will default to the value returned from "git config user.name"`,
	Run: func(cmd *cobra.Command, args []string) {
		filter := util.GetConfigOrExit("listfilter")
		filterReg := regexp.MustCompile(filter)
		r := util.OpenRepoOrExit()

		remote, err := r.Remote("origin")
		util.CheckSafeExit("Failed to get remote", err)

		items, err := remote.List(&git.ListOptions{})
		util.CheckSafeExit("Failed to list remote", err)

		fmt.Println("Remote Branches:")
		for _, item := range items {
			if item.Name().IsBranch() {
				commit, err := r.CommitObject(item.Hash())
				util.CheckSafeExit("Failed to get commit info", err)

				isAuthored := len(filterReg.FindAllStringIndex(commit.Author.Email, -1)) > 0
				if isAuthored {
					fmt.Println(item.Name().Short())
				}
			}
		}

		fmt.Println()
		fmt.Println("Local Branches:")
		refs, err := r.Branches()
		util.CheckSafeExit("Failed to get branches", err)

		refs.ForEach(func(ref *plumbing.Reference) error {
			commit, err := r.CommitObject(ref.Hash())
			util.CheckSafeExit("Failed to get commit info", err)

			isAuthored := len(filterReg.FindAllStringIndex(commit.Author.Email, -1)) > 0
			if isAuthored {
				fmt.Println(ref.Name().Short())
			}
			return nil
		})
	},
}

func newListCmd() *cobra.Command {
	return listCmd
}
