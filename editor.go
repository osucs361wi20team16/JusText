package justext

import "github.com/rivo/tview"

func EditorView() *tview.TextView {
	State.TextView = tview.NewTextView()
	State.TextView.Write([]byte("Hello, World!"))
	return State.TextView
}
