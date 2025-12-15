package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	Log func(msg string)
}

func main() {

	a := app.New()
	a.Settings().SetTheme(newMyTheme())

	w := a.NewWindow("Form Widget")
	w.Resize(fyne.NewSize(1200, 700))

	logArea := setLogArea()
	logArea.SetMinRowsVisible(5)
	scrollLog := container.NewVScroll(logArea)
	scrollLog.SetMinSize(fyne.NewSize(400, 200))

	cfg := Config{
		Log: func(msg string) {
			logArea.SetText(msg + "\n" + logArea.Text)
		},
	}

	form := cfg.setFilePathForm(w)
	rightSpacer := layout.NewSpacer()
	formContainer := container.NewBorder(
		nil, nil, nil, rightSpacer, form,
	)

	logAreaContainer := container.NewBorder(
		nil, nil, nil, rightSpacer,
		scrollLog,
	)

	mainContainer := container.NewVSplit(
		formContainer,
		logAreaContainer,
	)

	w.SetContent(mainContainer)
	w.ShowAndRun()

	log.Println("Exiting main")

}

func (cfg *Config) setFilePathForm(w fyne.Window) *widget.Form {
	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder("...")

	saveAsEntry := widget.NewEntry()
	now := time.Now()
	d2dinventoryname := fmt.Sprintf("Avis On Rent Inv (NATIONWIDE) %v", now.Format("01-02-06"))
	saveAsEntry.SetText(d2dinventoryname)

	pickBtn := widget.NewButton("Choose file...", func() {
		fd := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if r == nil {
				return // user cancelled
			}

			uri := r.URI()
			log.Println("Selected:", uri.Path())
			pathEntry.SetText(uri.Path())

			r.Close()
		}, w)

		fd.Show()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			widget.NewFormItem("File", pathEntry),
			widget.NewFormItem("Save As", saveAsEntry),
		},
		OnSubmit: func() {
			cfg.Log(fmt.Sprintf("Sending file (%v)...", pathEntry.Text))
			log.Println("Sending file...", pathEntry.Text)
			if err := cfg.sendSpreadsheetRequest(pathEntry.Text, saveAsEntry.Text); err != nil {
				log.Println("unable to generate spreadsheet:", err)
			}
		},
	}

	form.Append("", pickBtn)

	return form

}

func setLogArea() *widget.Entry {
	logEntry := widget.NewMultiLineEntry()
	return logEntry
}
