package justext

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView().SetDynamicColors(true)
	State.TextView.SetInputCapture(EditorInputCapture)
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	return State.TextView
}

func EditorInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyBS, tcell.KeyDEL:
		if len(State.Buffer) == 0 {
			return nil
		}
		State.Buffer = State.Buffer[:len(State.Buffer)-1]
		State.Cursor--
	case tcell.KeyEnter:
		State.Buffer = append(State.Buffer, '\n')
		State.Cursor++
	case tcell.KeyEscape:
		// Esc key on this level just passes focus to the Menu
		State.App.SetRoot(State.MainGrid, true)
		State.App.SetFocus(State.MenuGrid)
	default:
		State.Buffer = append(State.Buffer, byte(event.Rune()))
		State.Cursor++
	}
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	State.App.Draw()
	saveFile()
	return nil
}
func AddCursor(buffer []byte, cursor int) []byte {
	if cursor == len(buffer) {
		return append(buffer, []byte("[::r] [::-]")...)
	}

	return buffer
}
