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

// Select an item from a list
func Select(prompt string, options []string) string {
	return SelectHelp(prompt, "", options)
}

// SelectHelp select, with a help message
func SelectHelp(prompt string, help string, options []string) (val string) {
	val = ""
	question := &survey.Select{
		Message: prompt,
		Options: options,
	}
	if help != "" {
		question.Help = help
	}
	survey.AskOne(question, &val)
	return
}

// SelectMany options from a list
func SelectMany(prompt string, options []string) []string {
	return SelectManyHelp(prompt, "", options)
}

// SelectManyIndex returns the indicies of selected options from a list
func SelectManyIndex(prompt string, options []string) (val []int) {
	val = []int{}
	question := &survey.MultiSelect{
		Message: prompt,
		Options: options,
	}
	survey.AskOne(question, &val)
	return
}

// SelectManyHelp select many, but with a help prompt
func SelectManyHelp(prompt string, help string, options []string) (val []string) {
	val = []string{}
	question := &survey.MultiSelect{
		Message: prompt,
		Options: options,
	}
	if help != "" {
		question.Help = help
	}
	survey.AskOne(question, &val)
	return
}
