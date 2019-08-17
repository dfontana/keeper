package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Bool from the user for a yes no answer
func Bool(prompt string) (ans bool) {
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

// String will ask for a string response from the user, trimmed
func String(prompt string) string {
	fmt.Printf("%s ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Failed to scan input")
		os.Exit(1)
	}

	return strings.TrimSpace(scanner.Text())
}
