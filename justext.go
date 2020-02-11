package justext

import "github.com/rivo/tview"

type EditorState struct {
	Buffer   []rune
	App      *tview.Application
	TextView *tview.TextView
	Cursor   int
	Filename string
}

var State EditorState

func Run() {
	State = EditorState{
		Filename: "test.txt",
	}

	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetBorders(true).
		AddItem(MenuBarView(), 0, 0, 1, 1, 1, 1, false).
		AddItem(EditorView(), 1, 0, 1, 1, 1, 1, true).
		AddItem(StatusBarView(), 2, 0, 1, 1, 1, 1, false)

	State.App = tview.NewApplication()

	if err := State.App.
		SetRoot(grid, true).
		SetFocus(State.TextView).
		Run(); err != nil {
		panic(err)
	}
}
