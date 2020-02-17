package justext

import (
	"io/ioutil"
	"os"
	"log"
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
	fileReader, err := ioutil.ReadAll(file)

    if err != nil {
        fmt.Println("File reading error", err)
        return
	}

	State.Buffer = []byte(file)
    State.Filename = openFileName

    UpdateStatusBar("Editing " + "\"" + State.Filename + "\"!")
}
