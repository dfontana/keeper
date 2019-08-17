package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		filter := viper.GetString("listfilter")
		if filter == "" {
			fmt.Println("No list_filter found in ~/.keeper")
			os.Exit(1)
		}
		filterReg := regexp.MustCompile(filter)

		path, err := os.Getwd()
		if err != nil {
			fmt.Println("Failed to get working directory", err)
			os.Exit(0)
		}

		r, err := git.PlainOpen(path)
		if err != nil {
			fmt.Println("Failed to open repository", err)
			os.Exit(0)
		}

		remote, err := r.Remote("origin")
		if err != nil {
			fmt.Println("Failed to get remote", err)
			os.Exit(0)
		}

		items, err := remote.List(&git.ListOptions{})
		if err != nil {
			fmt.Println("Failed to list remote", err)
			os.Exit(0)
		}

		fmt.Println("Remote Branches:")
		for _, item := range items {
			if item.Name().IsBranch() {
				commit, err := r.CommitObject(item.Hash())
				if err != nil {
					fmt.Println("Failed to get commit info", err)
					os.Exit(0)
				}

				isAuthored := len(filterReg.FindAllStringIndex(commit.Author.Email, -1)) > 0
				if isAuthored {
					fmt.Println(item.Name().Short())
				}
			}
		}

		fmt.Println()
		fmt.Println("Local Branches:")
		refs, err := r.Branches()
		if err != nil {
			fmt.Println("Failed to get branches", err)
			os.Exit(0)
		}

		refs.ForEach(func(ref *plumbing.Reference) error {
			commit, err := r.CommitObject(ref.Hash())
			if err != nil {
				fmt.Println("Failed to get commit info", err)
				os.Exit(0)
			}
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
