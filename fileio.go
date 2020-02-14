package justext

import (
	"io/ioutil"
	"fmt"
)

func saveFile() bool {
	d1 := []byte(State.Buffer)
	err := ioutil.WriteFile(State.Filename, d1, 0700)
	if err != nil {
		panic(err)
	}
	return true
}

func openFile(openFileName string) {

	file, err := ioutil.ReadFile(openFileName)

    if err != nil {
        fmt.Println("File reading error", err)
        return
	}
	
	State.Buffer = []byte(file) 

	State.App.SetRoot(State.MainGrid, true)
	State.App.SetFocus(State.TextView)
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	State.App.Draw()
}
