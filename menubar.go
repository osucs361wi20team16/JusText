package justext

import (

	"os"
	"strconv"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// MenuBarView : ...
func MenuBarView() *tview.Grid {

	fileMenu := tview.NewDropDown().
		SetOptions([]string{"File", "Open", "Save", "Save As", "Quit"}, nil).
		SetCurrentOption(0)

	editMenu := tview.NewDropDown().
		SetOptions([]string{"Edit", "Word Count", "Copy", "Paste", "Select All"}, nil).
		SetCurrentOption(0)

	State.MenuGrid = tview.NewGrid().
		SetColumns(-1, -1, -1, -1, -1, -1).
		SetBorders(false).
		AddItem(fileMenu, 0, 0, 1, 1, 1, 1, true).
		AddItem(editMenu, 0, 1, 1, 1, 1, 1, false)

	fileMenu.SetSelectedFunc(func(text string, index int) {
		switch text {
		case "Open":
			path, _ := os.Getwd()
			table := listDir(path)
			runTable(table, State.App)

			//		form := tview.NewForm().
			//			AddInputField("File Name", "", 0, nil, nil)
			//		form.AddButton("Open", func() {
			//			openFile(form.GetFormItem(0).(*tview.InputField).GetText())
			//			DisplayEditor()
			//		})
			// State.App.SetRoot(form, true).SetFocus(form)
			fileMenu.SetCurrentOption(0)
		case "Save":
			saveFile()
			DisplayEditor()
			fileMenu.SetCurrentOption(0)
		case "Save As":
			form := tview.NewForm().
				AddInputField("File Name", "", 0, nil, nil)

			form.AddButton("Save", func() {
				State.Filename = form.GetFormItem(0).(*tview.InputField).GetText()
				saveFile()
				DisplayEditor()
			})

			State.App.SetRoot(form, true).SetFocus(form)
			fileMenu.SetCurrentOption(0)
		case "Quit":
			State.App.Stop()
		}
	})

	editMenu.SetSelectedFunc(func(text string, index int) {
		switch text {
		case "Word Count":
			UpdateStatusBar("Word count: " + strconv.Itoa(WordCount()))
		}
	})

	fileMenu.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			State.App.SetRoot(State.MainGrid, true)
			State.App.SetFocus(State.TextView)
			fileMenu.SetCurrentOption(0)
		}
		if key == tcell.KeyTab {
			State.App.SetFocus(editMenu)
			fileMenu.SetCurrentOption(0)
		}
	})

	editMenu.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			State.App.SetRoot(State.MainGrid, true)
			State.App.SetFocus(State.TextView)
			editMenu.SetCurrentOption(0)
		}
		if key == tcell.KeyTab {
			State.App.SetFocus(fileMenu)
			editMenu.SetCurrentOption(0)
		}
	})
	return State.MenuGrid
}
