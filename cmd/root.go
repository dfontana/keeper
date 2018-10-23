package cmd

import (
	"fmt"
	"os"

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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	initConfig()

	dirs := viper.GetStringMapString("dirs")
	for dir, abv := range dirs {
		rootCmd.PersistentFlags().BoolP(dir, abv, false, "Directory in codebase dir")
	}
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.SetConfigType("json")
	viper.SetConfigFile(home + "/.keeper")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
