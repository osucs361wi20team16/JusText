package justext

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func MenuBarView() *tview.Grid {

	fileMenu := tview.NewDropDown().
		SetOptions([]string{"File", "Open", "Save", "Save As", "Quit"}, nil).
		SetCurrentOption(0)

	editMenu := tview.NewDropDown().
		SetOptions([]string{"Edit", "Copy", "Paste", "Select All"}, nil).
		SetCurrentOption(0)

	State.MenuGrid = tview.NewGrid().
		SetColumns(-1, -1, -1, -1, -1, -1).
		SetBorders(false).
		AddItem(fileMenu, 0, 0, 1, 1, 1, 1, true).
		AddItem(editMenu, 0, 1, 1, 1, 1, 1, false)

	fileMenu.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			State.App.SetRoot(State.MainGrid, true)
			State.App.SetFocus(State.TextView)
		}
		if key == tcell.KeyTab {
			State.App.SetFocus(editMenu)
		}
	})

	editMenu.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			State.App.SetRoot(State.MainGrid, true)
			State.App.SetFocus(State.TextView)
		}
		if key == tcell.KeyTab {
			State.App.SetFocus(fileMenu)
		}
	})

	return State.MenuGrid
}
