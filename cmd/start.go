package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	gitConfig "gopkg.in/src-d/go-git.v4/config"
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

		template := viper.GetString("template")
		if template == "" {
			fmt.Println("No template found in ~/.keeper")
			os.Exit(1)
		}
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
		if err != nil {
			fmt.Println("Failed to open worktree")
			os.Exit(0)
		}

		status, err := worktree.Status()
		if err != nil {
			fmt.Println("Could not check dirtiness of branch")
			os.Exit(0)
		}

		if !status.IsClean() {
			fmt.Println("Dirty worktree, please commit or stash first")
			os.Exit(0)
		}

		// First try to checkout branch
		err = worktree.Checkout(&git.CheckoutOptions{Create: false, Force: false, Branch: b})

		if err != nil {
			// got an error  - try to create it
			err := worktree.Checkout(&git.CheckoutOptions{Create: true, Force: false, Branch: b})
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}

		templateRef := fmt.Sprintf("refs/heads/%s", template)
		testBranch := &gitConfig.Branch{
			Name:   template,
			Remote: "origin",
			Merge:  plumbing.ReferenceName(templateRef),
		}
		err = r.CreateBranch(testBranch)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	},
}

func newStartCmd() *cobra.Command {
	return startCmd
}
