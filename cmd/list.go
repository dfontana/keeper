package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var insensitive bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <search string>",
	Short: "List branches based on the search string",
	Long:  `Can search over the author name or branch name`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		filter := strings.Join(args, " ")
		params := []string{
			"git",
			"for-each-ref",
			"--format=' %(authorname) %09 %(refname)'",
			"--sort=authorname",
			"|",
			"grep",
			"--color=always",
			filter,
		}
		if insensitive {
			params = append(params, "-i")
		}

		gitExec := exec.Command("sh", "-c", strings.Join(params, " "))
		gitExec.Stdout = os.Stdout
		gitExec.Stderr = os.Stderr
		gitExec.Run()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(
		&insensitive,
		"insensitive",
		"i",
		false,
		"Help message for toggle",
	)
}
