package cmd

import (
	"fmt"
	"os"

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
		// var filter string
		// if len(args) == 0 {
		// 	params := []string{
		// 		"git",
		// 		"config",
		// 		"user.name",
		// 	}
		// 	out, err := util.Output(params)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	filter = string(out)
		// } else {
		// 	filter = strings.Join(args, " ")
		// }

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

		refs, err := r.Branches()
		if err != nil {
			fmt.Println("Failed to get branches", err)
			os.Exit(0)
		}

		refs.ForEach(func(ref *plumbing.Reference) error {
			fmt.Println(ref.Name())
			return nil
		})

		// params := []string{
		// 	"git",
		// 	"for-each-ref",
		// 	"--format=' %(authorname) %09 %(refname:short)'",
		// 	"--color=always",
		// 	"--sort=authorname",
		// 	"|",
		// 	"grep",
		// 	"--color=always",
		// 	"'" + string(filter) + "'",
		// }
		// if insensitive {
		// 	params = append(params, "-i")
		// }

		// util.Run(params)
	},
}

func newListCmd() *cobra.Command {
	listCmd.Flags().BoolVarP(
		&insensitive,
		"insensitive",
		"i",
		false,
		"Help message for toggle",
	)

	return listCmd
}
