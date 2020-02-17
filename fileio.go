package justext

import (
	"io/ioutil"
	"fmt"
)

func saveFile() bool {
	err := ioutil.WriteFile(State.Filename, State.Buffer, 0700)
	if err != nil {
		panic(err)
	}

    UpdateStatusBar("Saved to " + "\"" + State.Filename + "\"!")
	return true
}

func openFile(openFileName string) {

	file, err := ioutil.ReadFile(openFileName)

    if err != nil {
        fmt.Println("File reading error", err)
        return
	}

	State.Buffer = []byte(file)
    State.Filename = openFileName

    UpdateStatusBar("Editing " + "\"" + State.Filename + "\"!")
}
