package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
	"github.com/go-toast/toast"
)

var APPLICATION = app.NewWithID("kompass")
var WINDOW = APPLICATION.NewWindow("Kompass - Job Application Writing Tool")
var DelayInMs = 250 // Default delay in milliseconds per character
var text = ""       // Text to be written

func main() {
    WINDOW.Resize(fyne.NewSize(700, 350))
    WINDOW.SetFixedSize(true)

    iText := widget.NewMultiLineEntry()
    iText.Wrapping = fyne.TextWrapWord
    iText.SetPlaceHolder("Enter your text to write here...")

    iDelayInMs := widget.NewEntry()
    iDelayInMs.SetPlaceHolder("Enter delay in ms")

    iDelayContainer := container.NewWithoutLayout(iDelayInMs)
    iDelayInMs.Resize(fyne.NewSize(200, iDelayInMs.MinSize().Height)) // Set custom width

    bSubmit := widget.NewButton("Write", func() {
		// Write text then send notificaiton

		notification := toast.Notification{
			AppID:   "Kompass",
			Title:   "Kompass Notification",
			Message: "Writing complete",
		}

		notification.Push()
	})

    lCenter := container.NewMax(iText)
    lBottom := container.NewHBox(
        widget.NewLabel("Delay per character (ms):"),
        iDelayContainer,
        layout.NewSpacer(),
        bSubmit,
    )
    content := container.NewBorder(
        nil,
        lBottom,
        nil,
        nil,
        lCenter,
    )

    WINDOW.SetContent(content)
    WINDOW.ShowAndRun()
}