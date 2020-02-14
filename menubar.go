package justext

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func MenuBarView() *tview.Grid {

	fileMenu := tview.NewDropDown().
		SetOptions([]string{"File","Open", "Save", "Save As", "Quit"}, nil).
		SetCurrentOption(0)

	editMenu := tview.NewDropDown().
		SetOptions([]string{"Edit", "Copy", "Paste", "Select All"}, nil).
		SetCurrentOption(0)

	State.MenuGrid = tview.NewGrid().
		SetColumns(-1, -1, -1, -1, -1, -1).
		SetBorders(false).
		AddItem(fileMenu, 0, 0, 1, 1, 1, 1, true).
		AddItem(editMenu, 0, 1, 1, 1, 1, 1, false)

	fileMenu.SetSelectedFunc(func(text string, index int) {
		
		if text == "Open" {

			form := tview.NewForm().
				AddInputField("File Name", "", 20, nil, nil)
			
			form.AddButton("Open", func() {openFile(form.GetFormItem(0).(*tview.InputField).GetText())})
			form.AddButton("Cancel", func() {
				State.App.SetRoot(State.MainGrid, true)
				State.App.SetFocus(State.TextView)
			})



			State.App.SetRoot(form, true).SetFocus(form)
		}
		
		if text == "Save" {
			saveFile()
		}
		
		if text == "Save As" {
		
			form := tview.NewForm().
					AddInputField("File Name", "", 0, nil, nil)
		
			form.AddButton("Save", func() {
				State.Filename = form.GetFormItem(0).(*tview.InputField).GetText()
				saveFile()
			
				State.App.SetRoot(State.MainGrid, true)
				State.App.SetFocus(State.TextView)
				State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
				State.App.Draw()
			})

			State.App.SetRoot(form, true).SetFocus(form)
		}

		if text == "Quit" {
			State.App.Stop()
		}
	})

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
