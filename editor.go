package justext

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func DisplayEditor() {
	State.App.SetRoot(State.MainGrid, true)
	State.App.SetFocus(State.TextView)
	UpdateEditor()
}

// Redraw the editor view.
func UpdateEditor() {
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	State.App.Draw()
}

func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView().SetDynamicColors(true)
	State.TextView.SetInputCapture(EditorInputCapture)
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	return State.TextView
}

func EditorInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	// TODO — Handle arrow key presses
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
	case tcell.KeyRune:
		State.Buffer = append(State.Buffer, byte(event.Rune()))
		State.Cursor++
	default:
		UpdateEditor()
	}

	return nil
}

func AddCursor(buffer []byte, cursor int) []byte {
	if cursor == len(buffer) {
		return append(buffer, []byte("[::r] [::-]")...)
	} else {
		//ex: testing, cursor=3
		//  -> tes[::r]t[::-]ing

		cursor = 5

		// Pull out character where cursor is positioned
		cursorByte := buffer[cursor]

		// cursorBytes := []byte("[::r]" + string(cursorByte) + "[::-]")
		cursorBytes := []byte("//" + string(cursorByte) + "//")

		newBuffer := buffer
		_ = append(newBuffer[:cursor], cursorBytes...)
		// newBuffer = append(newBuffer, []byte("testing")...)
		_ = append(newBuffer, buffer[cursor+1:]...)
		// newBuffer = append(newBuffer, buffer[cursor+1:]...)

		return newBuffer
	}
}
