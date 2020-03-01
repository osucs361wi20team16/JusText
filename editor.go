package justext

import (
	"bytes"
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
	State.TextView.SetText(string(DisplayBuffer(State.Buffer, State.Cursor, State.Highlight)))
	State.App.Draw()
}

// EditorView : ...
func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView().SetDynamicColors(true)
	State.TextView.SetInputCapture(EditorInputCapture)
	State.TextView.SetText(string(DisplayBuffer(State.Buffer, State.Cursor, State.Highlight)))
	return State.TextView
}

// EditorInputCapture : ...
func EditorInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyBS, tcell.KeyDEL:
		if State.Highlight != nil {
			first, last := State.Highlight.First, State.Highlight.Last
			if last < first {
				first, last = last, first
			}
			if last == len(State.Buffer) {
				State.Buffer = State.Buffer[:first]
			} else {
				State.Buffer = append(
					State.Buffer[:first],
					State.Buffer[last+1:]...,
				)
			}
			State.Cursor = first
			State.Highlight = nil
		} else {
			if State.Cursor == 0 {
				return nil
			}
			if State.Cursor == len(State.Buffer) {
				State.Buffer = State.Buffer[:len(State.Buffer)-1]
			} else {
				State.Buffer = removeBytes(State.Buffer, State.Cursor-1, State.Cursor)
			}
			State.Cursor--
		}
	case tcell.KeyEnter:
		if State.Cursor == len(State.Buffer) {
			State.Buffer = append(State.Buffer, '\n')
		} else {
			State.Buffer = insertBytes(State.Buffer, State.Cursor, []byte{'\n'})
		}
		State.Cursor++
	case tcell.KeyEscape:
		if State.Highlight != nil {
			// First press clears highlight if it exists
			State.Highlight = nil
			break
		}
		// Esc key on this level just passes focus to the Menu
		State.App.SetRoot(State.MainGrid, true)
		State.App.SetFocus(State.MenuGrid)
	case tcell.KeyLeft:
		if State.Cursor == 0 {
			return nil
		}
		State.Cursor--
		if event.Modifiers() == tcell.ModShift {
			if State.Highlight == nil {
				State.Highlight = &Range{
					First: State.Cursor + 1,
					Last:  State.Cursor,
				}
			} else {
				State.Highlight.Last = State.Cursor
			}
		} else {
			State.Highlight = nil
		}
	case tcell.KeyRight:
		if State.Cursor == len(State.Buffer) {
			return nil
		}
		State.Cursor++
		if event.Modifiers() == tcell.ModShift {
			if State.Highlight == nil {
				State.Highlight = &Range{
					First: State.Cursor - 1,
					Last:  State.Cursor,
				}
			} else {
				State.Highlight.Last = State.Cursor
			}
		} else {
			State.Highlight = nil
		}
	case tcell.KeyUp:
		if State.Cursor == 0 {
			return nil
		}
		oldCursor := State.Cursor
		State.Cursor = findPrevLine(State.Buffer, State.Cursor)
		if event.Modifiers() == tcell.ModShift {
			if State.Highlight == nil {
				State.Highlight = &Range{
					First: oldCursor,
					Last:  State.Cursor,
				}
			} else {
				State.Highlight.Last = State.Cursor
			}
		} else {
			State.Highlight = nil
		}
	case tcell.KeyDown:
		if State.Cursor == len(State.Buffer) {
			return nil
		}
		oldCursor := State.Cursor
		State.Cursor = findNextLine(State.Buffer, State.Cursor)
		if event.Modifiers() == tcell.ModShift {
			if State.Highlight == nil {
				State.Highlight = &Range{
					First: oldCursor,
					Last:  State.Cursor,
				}
			} else {
				State.Highlight.Last = State.Cursor
			}
		} else {
			State.Highlight = nil
		}
	case tcell.KeyRune:
		if State.Highlight != nil {
			first, last := State.Highlight.First, State.Highlight.Last
			if last < first {
				first, last = last, first
			}
			if last == len(State.Buffer) {
				State.Buffer = State.Buffer[:first]
			} else {
				State.Buffer = append(
					append(
						State.Buffer[:first],
						byte(event.Rune()),
					),
					State.Buffer[last+1:]...,
				)
			}
			State.Cursor = first
			State.Highlight = nil
		} else {
			if State.Cursor == len(State.Buffer) {
				State.Buffer = append(State.Buffer, byte(event.Rune()))
			} else {
				State.Buffer = insertBytes(State.Buffer, State.Cursor, []byte{byte(event.Rune())})
			}
			State.Cursor++
		}
	default:
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

// DisplayBuffer : ...
func DisplayBuffer(buf []byte, cursor int, highlight *Range) []byte {
	displayBuffer := append(buf, ' ')
	var rng Range
	if highlight != nil {
		first, last := highlight.First, highlight.Last
		if last < first {
			first, last = last, first
		}
		rng = Range{
			First: first,
			Last:  last,
		}
	} else {
		rng = Range{
			First: cursor,
			Last:  cursor,
		}
	}
	l := displayBuffer[:rng.First]
	c := displayBuffer[rng.First : rng.Last+1]
	r := displayBuffer[rng.Last+1:]
	if bytes.Equal(c, []byte{'\n'}) {
		c = []byte("[::r] [::-]\n")
	} else {
		c = append(append([]byte("[::r]"), c...), []byte("[::-]")...)
	}
	parts := append([][]byte{}, l, c, r)
	displayBuffer = bytes.Join(parts, []byte{})
	if State.StatusBar != nil {
		UpdateStatusBar("\"" + string(displayBuffer) + "\"")
	}
	return displayBuffer
}

// WordCount : ...
// Original function source: https://www.dotnetperls.com/word-count-go
func WordCount() int {
	// Match non-space character sequences.
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	results := re.FindAllString(string(State.Buffer), -1)
	return len(results)
}
