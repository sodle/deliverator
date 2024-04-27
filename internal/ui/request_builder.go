package ui

import (
	"deliverator/internal/request"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func newRequestBuilder(request *request.Request) *fyne.Container {
	name := binding.BindString(&request.Name)
	method := binding.BindString(&request.Method)
	url := binding.BindString(&request.Url)

	nameField := widget.NewEntryWithData(name)

	methodField := &widget.Select{
		Options: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodHead,
			http.MethodTrace,
			http.MethodConnect,
			http.MethodOptions,
		},
		Selected:  request.Method,
		OnChanged: func(s string) { method.Set(s) },
	}

	urlField := widget.NewEntryWithData(url)

  output := widget.NewLabel("")

	rb := container.NewBorder(
		container.NewVBox(
			nameField,
			container.NewBorder(
				nil, nil,
				methodField, nil,
				urlField,
			),
			widget.NewButton("Send", func() {
        go func() {
          response, _ := request.Send()
          body, _ := io.ReadAll(response.Body)
          output.SetText(fmt.Sprintf("Status: %s\n\nBody:\n%s", response.Status, body))
        }()
      }),
		),
		nil, nil, nil,
    container.NewScroll(output),
	)

	return rb
}
