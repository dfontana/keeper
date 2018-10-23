package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <description> <number>",
	Short: "Begins a new brach with the given ticket number",
	Long: `Given the name and number will checkout a new branch from master with the given
	pt number in the properly named format`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			return
		}

		params := []string{
			"git",
			"checkout",
			"-b",
			fmt.Sprintf("%s_pt_%s", args[0], args[1]),
		}

		gitExec := exec.Command("sh", "-c", strings.Join(params, " "))
		gitExec.Stdout = os.Stdout
		gitExec.Stderr = os.Stderr
		gitExec.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
