package cmd

import (
	"fmt"
	"regexp"

	"github.com/dfontana/keeper/util"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// LocatedRef is a plumbing ref, but noted as local or remote
type LocatedRef struct {
	*plumbing.Reference
	isLocal bool
}

// IsLocal if the ref originates from the local repo or a remote
func (ref LocatedRef) IsLocal() bool {
	return ref.isLocal
}

// LocationName returns the short string of local or remote, depending on IsLocal
func (ref LocatedRef) LocationName() string {
	if ref.IsLocal() {
		return "local"
	}
	return "remote"
}

func listBranches(r *git.Repository, filter string) []*LocatedRef {
	filterReg := regexp.MustCompile(filter)
	branchRefs := []*LocatedRef{}

	// Get the remote branches
	remote, err := r.Remote("origin")
	util.CheckSafeExit("Failed to get remote", err)
	remoteRefs, err := remote.List(&git.ListOptions{})
	util.CheckSafeExit("Failed to list remote", err)
	for _, ref := range remoteRefs {
		if ref.Name().IsBranch() {
			commit, err := r.CommitObject(ref.Hash())
			util.CheckSafeExit("Failed to get commit info", err)
			isAuthored := len(filterReg.FindAllStringIndex(commit.Author.Email, -1)) > 0
			if isAuthored {
				branchRefs = append(branchRefs, &LocatedRef{ref, false})
			}
		}
	}

	// Get the local branches
	localRefs, err := r.Branches()
	util.CheckSafeExit("Failed to get branches", err)
	localRefs.ForEach(func(ref *plumbing.Reference) error {
		commit, err := r.CommitObject(ref.Hash())
		util.CheckSafeExit("Failed to get commit info", err)
		isAuthored := len(filterReg.FindAllStringIndex(commit.Author.Email, -1)) > 0
		if isAuthored {
			branchRefs = append(branchRefs, &LocatedRef{ref, true})
		}
		return nil
	})

	return branchRefs
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <search string>",
	Short: "List branches based on the search string",
	Long:  `Can search over the author name or branch name. If no search string is given, then this will default to the value returned from "git config user.name"`,
	Run: func(cmd *cobra.Command, args []string) {
		filter := util.GetConfigOrExit("listfilter")
		r := util.OpenRepoOrExit()
		branches := listBranches(r, filter)
		for _, branch := range branches {
			fmt.Println(fmt.Sprintf("%s\t%s", branch.LocationName(), branch.Name().Short()))
		}
	},
}

func newListCmd() *cobra.Command {
	return listCmd
}
