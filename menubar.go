package justext

import (
	"github.com/rivo/tview"
)

func MenuBarView() *tview.Grid {

	if State.menuGrid == nil {  // set up the drop down menus if they don't already exist
		fileMenu := tview.NewDropDown().
			SetOptions([]string{"File", "Open", "Save", "Save As", "Quit"}, nil).
			SetCurrentOption(0)

		editMenu := tview.NewDropDown().
			SetOptions([]string{"Edit", "Copy", "Paste", "Select All"}, nil).
			SetCurrentOption(0)

		if State.SwitchMenuColumn == true {
			State.menuGrid = tview.NewGrid().
			SetColumns(-1, -1, -1, -1, -1, -1).
			SetBorders(false).
			AddItem(editMenu, 0, 1, 1, 1, 1, 1, true).
			AddItem(fileMenu, 0, 0, 1, 1, 1, 1, false)
		} else if State.SwitchMenuColumn == false {
			State.menuGrid = tview.NewGrid().
			SetColumns(-1, -1, -1, -1, -1, -1).
			SetBorders(false).
			AddItem(editMenu, 0, 1, 1, 1, 1, 1, false).
			AddItem(fileMenu, 0, 0, 1, 1, 1, 1, true)
		}
		
	} else {  // otherwise, capture key strokes to change to other dropdown items or return focus to editor
		
		// func (g *Grid) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive))
		
	}

	return State.menuGrid
}

