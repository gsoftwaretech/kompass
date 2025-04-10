package main

import (
    // "github.com/go-vgo/robotgo"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    // "fyne.io/fyne/v2/widget"
    // "fyne.io/fyne/v2/container"
)

var APPLICATION = app.NewWithID("kompass")
var WINDOW = APPLICATION.NewWindow("Kompass - Job Application Writing Tool")

func main() {
    WINDOW.Resize(fyne.NewSize(700, 350))
    WINDOW.ShowAndRun()
}
