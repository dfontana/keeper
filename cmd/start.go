package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

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

		ack := util.PromptBool(fmt.Sprintf("Checkout to %s?", template))
		if !ack {
			fmt.Println("Cancelled.")
			os.Exit(0)
		}

		params := []string{
			"git",
			"checkout",
			"-b",
			template,
		}

		util.Run(params)
	},
}

func newStartCmd() *cobra.Command {
	return startCmd
}
