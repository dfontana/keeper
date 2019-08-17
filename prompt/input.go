package prompt

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Bool from the user for a yes no answer
func Bool(prompt string) bool {
	return BoolHelp(prompt, "")
}

// BoolHelp includes a help message
func BoolHelp(prompt string, help string) (val bool) {
	val = false
	question := &survey.Confirm{Message: prompt}
	if help != "" {
		question.Help = help
	}
	survey.AskOne(question, &val)
	return
}

// String will ask for a string response from the user, trimmed
func String(prompt string) string {
	return StringHelp(prompt, "")
}

// StringHelp includes a help message
func StringHelp(prompt string, help string) (val string) {
	val = ""
	question := &survey.Input{Message: prompt}
	if help != "" {
		question.Help = help
	}
	survey.AskOne(question, &val)
	val = strings.TrimSpace(val)
	return
}
