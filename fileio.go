package justext

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func saveFile() bool {
	err := ioutil.WriteFile(State.Filename, State.Buffer, 0666)
	if err != nil {
		panic(err)
	}

	UpdateStatusBar("Saved to " + "\"" + State.Filename + "\"!")
	return true
}

func openFile(openFileName string) {

	file, err := os.OpenFile(openFileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	State.Filename = openFileName
	fileReader, readError := ioutil.ReadAll(file)

	if readError != nil {
		fmt.Println("File reading error", err)
		return
	}

	State.Buffer = []byte(fileReader)
	State.Filename = openFileName
	UpdateEditor()

	UpdateStatusBar("Editing " + "\"" + State.Filename + "\"!")
}

func getFileName() {
	form := tview.NewForm().
		AddInputField("File Name", "", 0, nil, nil)

	form.AddButton("Save", func() {
		State.Filename = form.GetFormItem(0).(*tview.InputField).GetText()
		saveFile()
		DisplayEditor()
	})

	State.App.SetRoot(form, true).SetFocus(form).Run() // not sure why Run() is needed. Investigate.
}

func listDir(dir string, saveBool bool) *tview.Table {
	listTable := tview.NewTable()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	listTable.SetCell(0, 0, tview.NewTableCell("..").SetTextColor(tcell.ColorDimGray))
	rowCount, colCount := 1, 0
	if saveBool {
		listTable.SetCell(1, 0, tview.NewTableCell("<New File>").SetTextColor(tcell.ColorGreen))
		rowCount++
	}

	for _, f := range files {
		if f.IsDir() {
			listTable.SetCell(rowCount, colCount, tview.NewTableCell(f.Name()).SetTextColor(tcell.ColorDimGray))
		} else {
			listTable.SetCell(rowCount, colCount, tview.NewTableCell(f.Name()))
		}
		if rowCount > 20 {
			rowCount = 0
			colCount++
		} else {
			rowCount++
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	return listTable
}

func runTable(table *tview.Table, app *tview.Application, saveBool bool) {
	State.App.SetRoot(table, false).SetFocus(table)

	if !saveBool {
		table.SetBorder(true).SetTitle("Open File").SetTitleColor(tcell.ColorBlue)

	} else {
		table.SetBorder(true).SetTitle("Save File As").SetTitleColor(tcell.ColorRed)
	}

	table.Select(0, 0).SetFixed(1, 1).SetSelectable(true, true).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			DisplayEditor()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		if table.GetCell(row, column).Color == tcell.ColorDimGray { // this color indicates it is a directory

			path, _ := os.Getwd()
			dirStr := path + "/" + table.GetCell(row, column).Text
			os.Chdir(dirStr)
			if saveBool {
				table = listDir(dirStr, true)
			} else {
				table = listDir(dirStr, false)
			}
			runTable(table, app, saveBool)

		} else { // otherwise it is a file
			path, _ := os.Getwd()
			fileStr := path + "/" + table.GetCell(row, column).Text
			if saveBool {
				State.Filename = fileStr
				if table.GetCell(row, column).Color == tcell.ColorGreen {

					getFileName()

				} else {
					saveFile()
				}

			} else {
				openFile(fileStr)
			}
			DisplayEditor()
		}
	})
}
