package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

type Binding struct {
	ViewName    string
	Handler     func(*gocui.Gui, *gocui.View) error
	Key         interface{} // FIXME: find out how to get `gocui.Key | rune`
	Modifier    gocui.Modifier
	Description string
	Alternative string
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Mouse = true
	g.Cursor = true
	g.Highlight = true
	g.FgColor = gocui.ColorWhite
	g.SelFgColor = gocui.ColorCyan
	g.SelBgColor = gocui.ColorBlack
	g.BgColor = gocui.ColorBlack

	g.SetManager(gocui.ManagerFunc(layout), gocui.ManagerFunc(getFocusLayout()))

	setKeys(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func handleBranchSelect(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(v.Name())
	return err
}

func handleOuputSelect(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(v.Name())
	return err
}

func handleOpenCmdline(g *gocui.Gui, v *gocui.View) error {
	g.SetViewOnTop("cmdline")
	_, err := g.SetCurrentView("cmdline")
	return err
}

func handleCloseCmdline(g *gocui.Gui, v *gocui.View) error {
	g.SetViewOnBottom("cmdline")
	_, err := g.SetCurrentView("output")
	return err
}

func setKeys(g *gocui.Gui) error {
	listPanelMap := map[string]struct {
		prevLine func(*gocui.Gui, *gocui.View) error
		nextLine func(*gocui.Gui, *gocui.View) error
		focus    func(*gocui.Gui, *gocui.View) error
	}{
		"branches": {focus: handleBranchSelect}, // prevLine: gui.handleMenuPrevLine, nextLine: gui.handleMenuNextLine,
		"output":   {focus: handleOuputSelect},
	}

	var bindings []*Binding
	for viewName, functions := range listPanelMap {
		bindings = append(bindings, []*Binding{
			// {ViewName: viewName, Key: gocui.KeyArrowUp, Modifier: gocui.ModNone, Handler: functions.prevLine},
			// {ViewName: viewName, Key: gocui.KeyArrowDown, Modifier: gocui.ModNone, Handler: functions.nextLine},
			{ViewName: viewName, Key: gocui.MouseLeft, Modifier: gocui.ModNone, Handler: functions.focus},
		}...)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlE, gocui.ModNone, handleOpenCmdline); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, handleCloseCmdline); err != nil {
		return err
	}
	for _, binding := range bindings {
		if err := g.SetKeybinding(binding.ViewName, binding.Key, binding.Modifier, binding.Handler); err != nil {
			return err
		}
	}

	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("branches", 0, 0, int(0.15*float32(maxX)), maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Branches"
	}

	if v, err := g.SetView("output", int(0.15*float32(maxX))+1, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Output"
	}

	if v, err := g.SetView("cmdline", maxX/2-maxX/4, maxY/2, maxX/2+maxX/4, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetViewOnBottom("cmdline")
		v.Title = "Input Order"
		v.FgColor = gocui.ColorWhite
		v.Editable = true
		v.Wrap = true
	}

	return nil
}

// getFocusLayout returns a manager function for when view gain and lose focus
func getFocusLayout() func(g *gocui.Gui) error {
	var previousView *gocui.View
	return func(g *gocui.Gui) error {
		newView := g.CurrentView()
		if err := onFocusChange(g); err != nil {
			return err
		}
		if newView != previousView {
			if err := onFocusLost(previousView, newView); err != nil {
				return err
			}
			if err := onFocus(newView); err != nil {
				return err
			}
			previousView = newView
		}
		return nil
	}
}

func onFocusChange(g *gocui.Gui) error {
	currentView := g.CurrentView()
	for _, view := range g.Views() {
		view.Highlight = view == currentView
	}
	return nil
}

func onFocusLost(v *gocui.View, newView *gocui.View) error {
	return nil
}

func onFocus(v *gocui.View) error {
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
