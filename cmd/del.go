package cmd

import (
	"os"
	"os/exec"
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

		gitLocalExec := exec.Command(
			"sh",
			"-c",
			strings.Join(append(local, args...), " "),
		)

		gitRemoteExec := exec.Command(
			"sh",
			"-c",
			strings.Join(append(remote, args...), " "),
		)

		gitLocalExec.Stdout = os.Stdout
		gitLocalExec.Stderr = os.Stderr
		gitRemoteExec.Stdout = os.Stdout
		gitRemoteExec.Stderr = os.Stderr
		if err := gitLocalExec.Run(); err != nil {
			return
		}

		if err := gitRemoteExec.Run(); err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
