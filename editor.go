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
		if State.Cursor == 0 {
			return nil
		}
        if State.Cursor == len(State.Buffer) {
		    State.Buffer = State.Buffer[:len(State.Buffer)-1]
        } else {
            State.Buffer = removeBytes(State.Buffer, State.Cursor - 1, State.Cursor)
        }
		State.Cursor--
	case tcell.KeyEnter:
		State.Buffer = append(State.Buffer, '\n')
		State.Cursor++
	case tcell.KeyEscape:
		// Esc key on this level just passes focus to the Menu
		State.App.SetRoot(State.MainGrid, true)
		State.App.SetFocus(State.MenuGrid)
    case tcell.KeyLeft:
        if State.Cursor == 0 {
            return nil
        }
        State.Cursor--
    case tcell.KeyRight:
        if State.Cursor == len(State.Buffer) {
            return nil
        }
        State.Cursor++
	case tcell.KeyRune:
        if State.Cursor == len(State.Buffer) {
            State.Buffer = append(State.Buffer, byte(event.Rune()))
        } else {
            State.Buffer = insertBytes(State.Buffer, State.Cursor, []byte{byte(event.Rune())})
        }
		State.Cursor++
	default:
		UpdateEditor()
	}
    UpdateEditor()
	return nil
}

func insertBytes(buffer []byte, index int, newBytes []byte) []byte {

    firstHalf := make([]byte, index)
    copy(firstHalf, buffer[:index])
    secondHalf := make([]byte, len(buffer) - index)
    copy(secondHalf, buffer[index:])

    newBuffer := append(firstHalf, newBytes...)
    newBuffer = append(newBuffer, secondHalf...)

    return newBuffer
}

func removeBytes(buffer []byte, startId int, endId int) []byte {
    firstHalf := make([]byte, startId)
    copy(firstHalf, buffer[:startId])
    secondHalf := make([]byte, len(buffer) - endId)
    copy(secondHalf, buffer[endId:])

    newBuffer := append(firstHalf, secondHalf...)

    return newBuffer
}

func AddCursor(buffer []byte, cursor int) []byte {
	if cursor == len(buffer) {
		return append(buffer, []byte("[::r] [::-]")...)
	} else {
		// Pull out character where cursor is positioned
		cursorByte := buffer[cursor]

		cursorBytes := []byte("[::r]" + string(cursorByte) + "[::-]")

		firstHalf := make([]byte, cursor)
        copy(firstHalf, buffer[:cursor])
        secondHalf := make([]byte, len(buffer) - cursor)
        copy(secondHalf, buffer[cursor + 1:])

		newBuffer := append(firstHalf, cursorBytes...)
		newBuffer = append(newBuffer, secondHalf...)

		return newBuffer
	}
}
