package justext

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView()
	State.TextView.SetInputCapture(EditorInputHandler)
	State.TextView.SetText(string(State.Buffer))
	return State.TextView
}

func EditorInputHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyBS, tcell.KeyDEL:
		if len(State.Buffer) == 0 {
			return nil
		}
		State.Buffer = State.Buffer[:len(State.Buffer)-1]
	case tcell.KeyEnter:
		State.Buffer = append(State.Buffer, '\n')
	default:
		State.Buffer = append(State.Buffer, byte(event.Rune()))
	}
	State.TextView.SetText(string(State.Buffer))
	State.App.Draw()
	saveFile()
	return nil
}
