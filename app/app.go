package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dfontana/keeper/cmd"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keeper",
	Short: "Codebase Manager",
	Long: `The unofficial, unsupported, and certainly unmaintained
	codebase manager`,
}

// Start adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Start() {
	cmd.Init(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	err := initConfig()
	if err != nil {
		fmt.Println("You're missing the ~/.keeper config! Use `keeper generate` to make one")
	}
}

func initConfig() (err error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.SetConfigType("json")
	viper.SetConfigFile(filepath.Join(home, ".keeper"))
	err = viper.ReadInConfig()
	return
}
