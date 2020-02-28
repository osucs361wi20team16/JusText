package justext

import (
	"os"

	"github.com/rivo/tview"
)

type EditorState struct {
	Buffer    []byte
	App       *tview.Application
	TextView  *tview.TextView
	StatusBar *tview.TextView
	Cursor    int
	MenuGrid  *tview.Grid
	MainGrid  *tview.Grid
	Filename  string
}

var State EditorState

func Run() {

	State.MainGrid = tview.NewGrid().
		SetRows(1, 0, 1).
		SetBorders(true).
		AddItem(MenuBarView(), 0, 0, 1, 1, 1, 1, false).
		AddItem(EditorView(), 1, 0, 1, 1, 1, 1, true).
		AddItem(StatusBarView(), 2, 0, 1, 1, 1, 1, false)

	State.App = tview.NewApplication()

	if len(os.Args) == 1 {
		State.Filename = "test.txt" // open this file by default
		openFile(State.Filename)
	} else {
		State.Filename = os.Args[1] // if there is a command line argument, open that instead
		openFile(State.Filename)
	}

	if err := State.App.
		SetRoot(State.MainGrid, true).
		SetFocus(State.TextView).
		Run(); err != nil {
		panic(err)
	}
}
