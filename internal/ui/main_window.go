package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"deliverator/internal/request"
)

func NewMainWindow(app fyne.App) fyne.Window {
	activeRequest := &request.Request{
		Name:   "Example Request",
		Url:    "https://example.com",
		Method: "GET",
	}
	w := app.NewWindow("Deliverator")
	w.Resize(fyne.NewSize(800, 600))

	w.SetContent(container.NewBorder(
		nil, nil, nil, nil,
		newRequestBuilder(activeRequest),
	))

	return w
}
