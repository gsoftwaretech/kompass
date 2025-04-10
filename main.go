package main

import (
    // "github.com/go-vgo/robotgo"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    // "fyne.io/fyne/v2/widget"
    // "fyne.io/fyne/v2/container"
	// "time"
)

var APPLICATION = app.NewWithID("kompass")
var WINDOW 		= APPLICATION.NewWindow("Kompass - Job Application Writing Tool")
var WPM         = 40  // Default general WPM
var delay       = 500 // Default delay of 500ms 
var text     	= "" // Text to be written

func main() {
    WINDOW.Resize(fyne.NewSize(700, 350))
	WINDOW.SetFixedSize(true)

    WINDOW.ShowAndRun()
}
