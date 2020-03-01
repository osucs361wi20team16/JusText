package justext

import (
	"bytes"
	"fmt"
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
		if State.Cursor == 0 {
			return nil
		}
		if State.Highlight != nil {
			State.Highlighting = false
			State.Buffer = append(
				State.Buffer[:State.Highlight.First],
				State.Buffer[State.Highlight.Last:]...,
			)
			State.Cursor = State.Highlight.First
			State.Highlight = nil
		} else {
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
			State.Highlighting = false
			State.Highlight = nil
			break
		}
		// Esc key on this level just passes focus to the Menu
		State.App.SetRoot(State.MainGrid, true)
		State.App.SetFocus(State.MenuGrid)
	case tcell.KeyCtrlSpace:
		if State.Highlighting {
			State.Highlighting = false
			break
		}
		State.Highlighting = true
		State.Highlight = &Range{
			First: State.Cursor,
			Last:  State.Cursor,
		}
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
	}
	if State.Highlighting {
		UpdateHighlight()
	}
	UpdateEditor()
	return nil
}

func UpdateHighlight() {
	State.Highlight.Last = State.Cursor
	UpdateStatusBar(fmt.Sprintf("%#v", State.Highlight))
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
	toReverse := []Range{}
	if highlight != nil {
		first := highlight.First
		last := highlight.Last
		if last < first {
			first, last = last, first
		}
		switch {
		case cursor < first:
			toReverse = append(
				toReverse,
				Range{First: cursor, Last: cursor},
				Range{First: first, Last: last},
			)
		case cursor > last:
			toReverse = append(
				toReverse,
				Range{First: first, Last: last},
				Range{First: cursor, Last: cursor},
			)
		case cursor == first:
			toReverse = append(
				toReverse,
				Range{First: first + 1, Last: last},
			)
		case cursor == last:
			toReverse = append(
				toReverse,
				Range{First: first, Last: last - 1},
			)
		case cursor > first && cursor < last:
			toReverse = append(
				toReverse,
				Range{First: first, Last: cursor - 1},
				Range{First: cursor + 1, Last: last},
			)
		}
	} else {
		toReverse = append(toReverse, Range{
			First: cursor,
			Last:  cursor,
		})
	}

	ix := 0
	parts := [][]byte{}
	for _, rng := range toReverse {
		l := displayBuffer[ix:rng.First]
		c := displayBuffer[rng.First : rng.Last+1]
		ix = rng.Last + 1
		if bytes.Equal(c, []byte{'\n'}) {
			c = []byte("[::r] [::-]\n")
		} else {
			c = append(append([]byte("[::r]"), c...), []byte("[::-]")...)
		}
		parts = append(parts, l, c)
	}
	parts = append(parts, displayBuffer[ix:])
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
