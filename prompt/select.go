package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// Option has a value it represents and whether that item is chosen.
type Option struct {
	Selected bool
	Value    string
}

// Select from a list of items repeatedly until enter is pressed.
// Using space will select
func Select(options []Option) []Option {

	// the questions to ask
	var qs = []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "What is your name?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "color",
			Prompt: &survey.Select{
				Message: "Choose a color:",
				Options: []string{"red", "blue", "green"},
				Default: "red",
			},
		},
		{
			Name:   "age",
			Prompt: &survey.Input{Message: "How old are you?"},
		},
	}

	answers := struct {
		Name          string // survey will match the question and field names
		FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
		Age           int    // if the types don't match, survey will convert it
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return nil
}
