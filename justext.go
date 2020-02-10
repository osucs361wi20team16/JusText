package justext

import "github.com/rivo/tview"

type EditorState struct {
	Buffer   string
	App      *tview.Application
	TextView *tview.TextView
	Cursor   int
	menuGrid *tview.Grid
	maingrid *tview.Grid
	Filename string
	SwitchMenuColumn bool
}

var State EditorState

func Run() {
	State = EditorState{}
	State.SwitchMenuColumn = false
	State.Filename = "test.txt"
	State.maingrid = tview.NewGrid().
		SetRows(1, 0, 1).
		SetBorders(true).
		AddItem(MenuBarView(), 0, 0, 1, 1, 1, 1, false).
		AddItem(EditorView(), 1, 0, 1, 1, 1, 1, true).
		AddItem(StatusBarView(), 2, 0, 1, 1, 1, 1, false)

	State.App = tview.NewApplication()

	if err := State.App.
		SetRoot(State.maingrid, true).
		SetFocus(State.TextView).
		Run(); err != nil {
		panic(err)
	}
}
