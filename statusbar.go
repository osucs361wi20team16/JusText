package justext

import "github.com/rivo/tview"

func StatusBarView() *tview.TextView {
	State.StatusBar = tview.NewTextView()
	State.StatusBar.SetText("JustText")
	return State.StatusBar
	// return tview.NewTextView().SetText("JusText")
}

func UpdateStatusBar(text string) {
	State.StatusBar.SetText(text)
	State.App.Draw()
}
