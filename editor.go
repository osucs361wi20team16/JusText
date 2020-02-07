package justext

import (
	"unicode/utf8"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView()
	State.TextView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		R := event.Rune()
		B := make([]byte, 1)
		utf8.EncodeRune(B, R)
		State.TextView.Write(B)
		State.App.Draw()
		return nil
	})
	State.TextView.Write([]byte("Hello, World!"))
	return State.TextView
}
