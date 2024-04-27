package main

import (
	"fyne.io/fyne/v2/app"

	"deliverator/internal/ui"
)

func main() {
	a := app.NewWithID("deliverator")
	w := ui.NewMainWindow(a)
  w.ShowAndRun()
}
