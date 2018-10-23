package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		// Determine command to run
		isFetch, _ := cmd.Flags().GetBool("fetch")
		if isFetch {
			fetchDirs(cmd, args)
		}
	},
}

func fetchDirs(cmd *cobra.Command, args []string) {
	executeDirs := []string{}
	dirs := viper.GetStringMapString("dirs")
	for k := range dirs {
		if ok, _ := cmd.Flags().GetBool(k); ok {
			executeDirs = append(executeDirs, k)
		}
	}

	// TODO continue from here
	fmt.Println(executeDirs)
}

func init() {
	rootCmd.AddCommand(eachCmd)

	// Prepare the commands
	eachCmd.Flags().Bool("fetch", false, "Fetches each of the flagged dirs")
}
