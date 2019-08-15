package util

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run will execute a command with its output to the command line
func Run(params []string) (err error) {
	cmd := exec.Command("sh", "-c", strings.Join(params, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

// RunString executes the given string command
func RunString(command string) (err error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

// Output returns the result of a command from the command line
func Output(params []string) (out string, err error) {
	cmd := exec.Command("sh", "-c", strings.Join(params, " "))
	res, err := cmd.Output()
	out = string(res)
	return
}

// BuildDirs inspects the given cmd and builds a list of absolute path working dirs
func BuildDirs(cmd *cobra.Command) (executeDirs []string) {
	// Get root, expanded if needed
	root := viper.GetString("codebase")
	if root[0] == '~' {
		root = root[1:]
	}
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Build dirs to run
	dirs := viper.GetStringMapString("dirs")
	for k := range dirs {
		if ok, _ := cmd.Flags().GetBool(k); ok {
			executeDirs = append(
				executeDirs,
				filepath.Join(usr.HomeDir, root, k),
			)
		}
	}
	return
}

// PromptString will ask for a string response from the user, trimmed
func PromptString(prompt string) (ans string) {
	fmt.Printf("%s ", prompt)
	_, err := fmt.Scanln(&ans)
	if err != nil {
		panic(err)
	}
	ans = strings.TrimSpace(ans)
	return
}

// PromptBool the user a yes no answer
func PromptBool(prompt string) (ans bool) {
	fmt.Printf("%s? [y/n]: ", prompt)
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	ans = s[0] == 'y'
	return
}
