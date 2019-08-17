package util

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
)

// RunString executes the given string command
func RunString(command string) (err error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

// ValidateStringSpaces to not contain spaces or is empty
func ValidateStringSpaces(value string) bool {
	space := regexp.MustCompile(" ")
	numSpaces := len(space.FindAllStringIndex(value, -1))
	return numSpaces == 0 || value == ""
}

// OpenRepoOrExit in the current working directory, or exit
func OpenRepoOrExit() *Repository {
	path, err := os.Getwd()
	CheckSafeExit("Failed to get working directory", err)

	r, err := git.PlainOpen(path)
	CheckSafeExit("Failed to open repository", err)
	return r
}

// CheckSafeExit if the error exists with message
func CheckSafeExit(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(0)
	}
}

// GetConfigOrExit from keeper config or exit program
func GetConfigOrExit(key string) string {
	val := viper.GetString(key)
	CheckSafeExit(fmt.Sprintf("No %s found in ~/.keeper", key))
	return val
}
