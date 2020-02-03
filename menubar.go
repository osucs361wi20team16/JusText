package justext

import "github.com/rivo/tview"

func MenuBarView() *tview.Table {
	menubar := tview.NewTable()

	fileMenu := tview.NewTableCell("File")
	editMenu := tview.NewTableCell("Edit")

	menubar.SetCell(0, 0, fileMenu)
	menubar.SetCell(0, 1, editMenu)

	return menubar
}
