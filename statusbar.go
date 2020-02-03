package justext

import "github.com/rivo/tview"

func StatusBarView() *tview.TextView {
	return tview.NewTextView().SetText("JusText")
}
