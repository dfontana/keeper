package prompt

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

// Option has a value it represents and whether that item is chosen.
type Option struct {
	Selected bool
	Value    string
}

// Select from a list of items repeatedly until enter is pressed.
// Using space will select
func Select(options []Option) []Option {

	// init screen
	encoding.Register()
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	s.EnableMouse()

	// Defer close
	defer s.Fini()

	// width, height := s.Size()
	return nil
}
