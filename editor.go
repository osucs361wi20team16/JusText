package justext

import (
	"strings"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	
)

func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView()
	State.TextView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyBS, tcell.KeyDEL:
			if len(State.Buffer) == 0 {
				break
			}
			RS := []rune(State.Buffer)
			RS = RS[:len(RS)-1]
			State.Buffer = string(RS)
		case tcell.KeyEscape:
			// Esc key on this level just passes focus to the Menu 
			State.App.SetRoot(State.maingrid, true)
			State.App.SetFocus(State.menuGrid)
		default:
			SB := &strings.Builder{}
			SB.WriteString(State.Buffer)
			SB.WriteRune(event.Rune())
			State.Buffer = SB.String()
		}
		State.TextView.SetText(State.Buffer)
		State.App.Draw()
		saveFile()
		return nil
	})
	State.TextView.SetText(State.Buffer)
	return State.TextView
}
