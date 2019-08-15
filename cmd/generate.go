package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/dfontana/keeper/util"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type config struct {
	Codebase string            `json:"codebase"`
	Dir      map[string]string `json:"dirs"`
}

// generateCmd represents the generate config command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Creates a new ~/.keeper config",
	Long:  `Creates the ~/.keeper config file the program uses to operate. Will replace existing one if present.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config = config{
			Codebase: "",
			Dir:      map[string]string{},
		}

		exists := false
		fmt.Println("First provide the root directory of your codebase, where all sub-repos are")
		for !exists {
			ans := strings.Trim(util.PromptString("Codebase Path:"), " ")
			ans, _ = homedir.Expand(ans)
			_, err := os.Stat(ans)
			if err == nil {
				config.Codebase = ans
				exists = true
			}
		}

		done := false
		exists = false
		fmt.Println("Now add directories and shorthand flag in codebase, ex: javascript j")
		for !done || !exists {
			var directory, flag string
			fmt.Printf("Entry: ")
			fmt.Scanf("%s %s", &directory, &flag)
			_, err := os.Stat(filepath.Join(config.Codebase, directory))
			if err == nil && strings.Trim(directory, " ") != "" && strings.Trim(flag, " ") != "" {
				config.Dir[directory] = flag
				exists = true
				done = !util.PromptBool("Add Another")
			}
		}

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		jsonFile, err := os.Create(filepath.Join(home, ".keeper"))
		if err != nil {
			fmt.Println("Error creating Config file:", err)
			return
		}
		defer jsonFile.Close()

		jsonWriter := io.Writer(jsonFile)
		encoder := json.NewEncoder(jsonWriter)
		err = encoder.Encode(&config)
		if err != nil {
			fmt.Println("Error encoding Config to file:", err)
			return
		}

		return
	},
}

func newGenerateCmd() *cobra.Command {
	return generateCmd
}
