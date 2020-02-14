package justext

import (
	"io/ioutil"
	"os"
	"log"
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
	
	file, err := os.OpenFile(openFileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	State.Filename = openFileName
	fileReader, err := ioutil.ReadAll(file)

	State.Buffer = []byte(fileReader) 
	

	State.App.SetRoot(State.MainGrid, true)
	State.App.SetFocus(State.TextView)
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	State.App.Draw()
}
