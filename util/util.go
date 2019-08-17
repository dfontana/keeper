package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

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

// PromptString will ask for a string response from the user, trimmed
func PromptString(prompt string) string {
	fmt.Printf("%s ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Failed to scan input")
		os.Exit(1)
	}

	return strings.TrimSpace(scanner.Text())
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
