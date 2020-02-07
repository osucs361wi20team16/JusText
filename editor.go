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
		case tcell.KeyBS:
			if len(State.Buffer) == 0 {
				break
			}
			RS := []rune(State.Buffer)
			RS = RS[:len(RS)-1]
			State.Buffer = string(RS)
		default:
			SB := &strings.Builder{}
			SB.WriteString(State.Buffer)
			SB.WriteRune(event.Rune())
			State.Buffer = SB.String()
		}
		State.TextView.SetText(State.Buffer)
		State.App.Draw()
		return nil
	})
	State.TextView.SetText("")
	return State.TextView
}
