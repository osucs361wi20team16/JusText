package justext

import (
	"io/ioutil"
)

func saveFile() bool {
	d1 := []byte(State.Buffer)
	err := ioutil.WriteFile(State.Filename, d1, 0700)
	if err != nil {
		panic(err)
	}
	return true
}
