package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

var fetch bool

// eachCmd represents the each command
var eachCmd = &cobra.Command{
	Use:   "each",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get Dirs
		dirs := BuildDirs(cmd)

		// Determine command to run
		isFetch, _ := cmd.Flags().GetBool("fetch")
		if isFetch {
			var wg sync.WaitGroup
			for _, dir := range dirs {
				wg.Add(1)
				go fetchDir(dir, &wg)
				wg.Wait()
			}
			return
		}
	},
}

// fetchDir fetches the given git directory
func fetchDir(dir string, wg *sync.WaitGroup) {
	defer wg.Done()

	params := []string{
		"git",
		"fetch",
		"origin",
		"master",
	}
	if err := Run(params); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(eachCmd)

	// Prepare the commands
	eachCmd.Flags().Bool("fetch", false, "Fetches each of the flagged dirs")
}
