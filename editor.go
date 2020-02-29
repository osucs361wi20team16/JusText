package justext

import (
	"regexp"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func DisplayEditor() {
	State.App.SetRoot(State.MainGrid, true)
	State.App.SetFocus(State.TextView)
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

	}
	State.TextView.SetText(string(AddCursor(State.Buffer, State.Cursor)))
	State.App.Draw()
	return nil
}
func AddCursor(buffer []byte, cursor int) []byte {
	if cursor == len(buffer) {
		return append(buffer, []byte("[::r] [::-]")...)
	}

	return buffer
}

// Original function source: https://www.dotnetperls.com/word-count-go

func WordCount() int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(string(State.Buffer), -1)
	return len(results)
}
