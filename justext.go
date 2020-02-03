package justext

import "github.com/rivo/tview"

func Run() {
	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetBorders(true).
		AddItem(MenuBarView(), 0, 0, 1, 1, 1, 1, false).
		AddItem(EditorView(), 1, 0, 1, 1, 1, 1, true).
		AddItem(StatusBarView(), 2, 0, 1, 1, 1, 1, false)

	if err := tview.NewApplication().SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
