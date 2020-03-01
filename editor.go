package justext

import (
	"regexp"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// DisplayEditor : Update editor and bring into focus
func DisplayEditor() {
	State.App.SetRoot(State.MainGrid, true)
	State.App.SetFocus(State.TextView)
	UpdateEditor()
}

// UpdateEditor : Redraw the editor view.
func UpdateEditor() {
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	State.App.Draw()
}

// EditorView : ...
func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView().SetDynamicColors(true)
	State.TextView.SetInputCapture(EditorInputCapture)
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	return State.TextView
}

// EditorInputCapture : ...
func EditorInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyBS, tcell.KeyDEL:
		if State.Cursor == 0 {
			return nil
		}
		if State.Cursor == len(State.Buffer) {
			State.Buffer = State.Buffer[:len(State.Buffer)-1]
		} else {
			State.Buffer = removeBytes(State.Buffer, State.Cursor-1, State.Cursor)
		}
		State.Cursor--
	case tcell.KeyEnter:
		if State.Cursor == len(State.Buffer) {
			State.Buffer = append(State.Buffer, '\n')
		} else {
			State.Buffer = insertBytes(State.Buffer, State.Cursor, []byte{'\n'})
		}
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
	case tcell.KeyUp:
		if State.Cursor == 0 {
			return nil
		}
		State.Cursor = findPrevLine(State.Buffer, State.Cursor)
	case tcell.KeyDown:
		if State.Cursor == len(State.Buffer) {
			return nil
		}
		State.Cursor = findNextLine(State.Buffer, State.Cursor)
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

func findPrevLine(buffer []byte, index int) int {
	distance := 1
	_, _, width, _ := State.TextView.GetInnerRect()
	for distance < width && distance < index &&
		buffer[index-distance] != '\n' {
		distance++
	}
	return index - distance
}

func findNextLine(buffer []byte, index int) int {
	distance := 1
	_, _, width, _ := State.TextView.GetInnerRect()
	for distance < width && distance+index < len(buffer)-1 &&
		buffer[index+distance] != '\n' {
		distance++
	}
	return index + distance
}

func insertBytes(buffer []byte, index int, newBytes []byte) []byte {
	firstHalf := make([]byte, index)
	copy(firstHalf, buffer[:index])
	secondHalf := make([]byte, len(buffer)-index)
	copy(secondHalf, buffer[index:])

	newBuffer := append(firstHalf, newBytes...)
	newBuffer = append(newBuffer, secondHalf...)

	return newBuffer
}

func removeBytes(buffer []byte, startID int, endID int) []byte {
	firstHalf := make([]byte, startID)
	copy(firstHalf, buffer[:startID])
	secondHalf := make([]byte, len(buffer)-endID)
	copy(secondHalf, buffer[endID:])

	newBuffer := append(firstHalf, secondHalf...)

	return newBuffer
}

// AddCursor : Creates a copy of the current buffer with the
// cursor inserted using tview bracket syntax.
func AddCursor(buffer []byte, cursor int) []byte {
	if cursor == len(buffer) {
		return append(buffer, []byte("[::r] [::-]")...)
	}

	// Pull out character where cursor is positioned
	cursorByte := buffer[cursor]

	var cursorBytes []byte
	if cursorByte == '\n' {
		cursorBytes = []byte("[::r] [::-]" + "\n")
	} else {
		cursorBytes = []byte("[::r]" + string(cursorByte) + "[::-]")
	}

	firstHalf := make([]byte, cursor)
	copy(firstHalf, buffer[:cursor])
	secondHalf := make([]byte, len(buffer)-cursor)
	copy(secondHalf, buffer[cursor+1:])

	newBuffer := append(firstHalf, cursorBytes...)
	newBuffer = append(newBuffer, secondHalf...)

	return newBuffer
}

// Original function source: https://www.dotnetperls.com/word-count-go

func WordCount() int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(string(State.Buffer), -1)
	return len(results)
}
