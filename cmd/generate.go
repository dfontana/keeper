package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dfontana/keeper/prompt"
	"github.com/dfontana/keeper/util"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type keeperConfig struct {
	Namespace  string   `json:"namespace"`
	ListFilter string   `json:"listfilter"`
	Template   string   `json:"template"`
	Prompts    []string `json:"prompts"`
}

// generateCmd represents the generate config command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Creates a new ~/.keeper config",
	Long:  `Creates the ~/.keeper config file the program uses to operate. Will replace existing one if present.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config = keeperConfig{
			Namespace: "",
			Template:  "",
			Prompts:   []string{},
		}

		valid := false
		for !valid {
			config.Namespace = prompt.String("What's your K8s Namespace")
			if util.ValidateStringSpaces(config.Namespace) {
				valid = true
			} else {
				fmt.Println("Value cannot be empty or contain spaces")
			}
		}

		valid = false
		for !valid {
			config.ListFilter = prompt.StringHelp(
				"Supply the default filter for list",
				"This will filter over commiter emails",
			)
			if config.ListFilter != "" {
				valid = true
			} else {
				fmt.Println("Value cannot be empty")
			}
		}

		valid = false
		for !valid {
			config.Template = prompt.StringHelp(
				"Provide the template for new branches",
				"Specify #s# where you want to prompt for input. You'll then specify these prompts afterwards.",
			)
			if util.ValidateStringSpaces(config.Template) {
				valid = true
			} else {
				fmt.Println("Value cannot be empty or contain spaces")
			}
		}

		pcntS := regexp.MustCompile("#s#")
		numPrompts := len(pcntS.FindAllStringIndex(config.Template, -1))

		valid = false || numPrompts == 0
		if !valid {
			fmt.Println("Provide your prompts for each placeholder: ")
		}
		for !valid || numPrompts != 0 {
			nextPrompt := prompt.String("Prompt:")
			if nextPrompt != "" {
				valid = true
				numPrompts--
				config.Prompts = append(config.Prompts, nextPrompt)
			} else {
				fmt.Println("Value cannot be empty")
			}
		}

		home, err := homedir.Dir()
		util.CheckSafeExit("Can't find homedir", err)

		jsonFile, err := os.Create(filepath.Join(home, ".keeper"))
		util.CheckSafeExit("Error creating Config file:", err)
		defer jsonFile.Close()

		jsonWriter := io.Writer(jsonFile)
		encoder := json.NewEncoder(jsonWriter)
		err = encoder.Encode(&config)
		util.CheckSafeExit("Error encoding Config to file:", err)

		fmt.Println("Config written.")
		return
	},
}

func newGenerateCmd() *cobra.Command {
	return generateCmd
}
