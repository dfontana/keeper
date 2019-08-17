package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"

	"github.com/dfontana/keeper/util"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Begin a new branch",
	Long: `Using the template in your config file, beings a new branch
	after prompting for template values`,
	Run: func(cmd *cobra.Command, args []string) {
		r := util.OpenRepoOrExit()

		valid := false
		prompts := viper.GetStringSlice("prompts")
		values := []string{}
		for _, prompt := range prompts {
			for !valid {
				value := util.PromptString(fmt.Sprintf("%s:", prompt))
				if util.ValidateStringSpaces(value) {
					valid = true
					values = append(values, value)
				} else {
					fmt.Println("Value may not have spaces or be empty")
				}
			}
			valid = false
		}

		template := util.GetConfigOrExit("template")
		for _, value := range values {
			template = strings.Replace(template, "#s#", value, 1)
		}

		ack := util.PromptBool(fmt.Sprintf("Checkout to %s", template))
		if !ack {
			fmt.Println("Cancelled.")
			os.Exit(0)
		}

		branch := fmt.Sprintf("refs/heads/%s", template)
		b := plumbing.ReferenceName(branch)
		worktree, err := r.Worktree()
		util.CheckSafeExit("Failed to open worktree", err)

		status, err := worktree.Status()
		util.CheckSafeExit("Could not check dirtiness of branch", err)

		if !status.IsClean() {
			fmt.Println("Dirty worktree, please commit or stash first")
			os.Exit(0)
		}

		err = worktree.Checkout(&git.CheckoutOptions{Create: false, Force: false, Branch: b})
		if err != nil {
			err := worktree.Checkout(&git.CheckoutOptions{Create: true, Force: false, Branch: b})
			util.CheckSafeExit("Can't checkout branch", err)
		}
	},
}

func newStartCmd() *cobra.Command {
	return startCmd
}
