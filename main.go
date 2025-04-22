package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
	"github.com/go-toast/toast"
    "github.com/go-vgo/robotgo"
    "strconv"
)

var APPLICATION         = app.NewWithID("kompass")
var WINDOW              = APPLICATION.NewWindow("Kompass - Writing Tool")
var DisplayNotification = true // Flag to display notification

func main() {
    WINDOW.Resize(fyne.NewSize(700, 350))
    WINDOW.SetFixedSize(true)

    iText := widget.NewMultiLineEntry()
    iText.Wrapping = fyne.TextWrapWord
    iText.SetPlaceHolder("Enter your text to write here...")

    iDelayInMs := widget.NewEntry()
    iDelayInMs.SetPlaceHolder("Enter delay in ms")

    iDelayContainer := container.NewWithoutLayout(iDelayInMs)
    iDelayInMs.Resize(fyne.NewSize(160, iDelayInMs.MinSize().Height)) // Set custom width

    iNotificationCheck := widget.NewCheck("", func(checked bool) {
        DisplayNotification = checked
    })
    iNotificationCheck.SetChecked(true) // Default to checked

    StopWriting := false // Flag to stop writing
    bStop   := widget.NewButton("Stop", func() {
        StopWriting = true
    })

    bSubmit := widget.NewButton("Write", func() {
        StopWriting = false // Reset the stop flag
		// Write text then send notificaiton

        text       := iText.Text
        delay, err := strconv.Atoi(iDelayInMs.Text)
        if (err != nil) {
            delay = 25 // Default delay of 25ms
        }

        go func() {
            // Set 2s delay
            robotgo.MilliSleep(2000)
            
            for _, char := range text {
                if StopWriting {
                    if (DisplayNotification) {
                        notification := toast.Notification{
                            AppID:   "Kompass",
                            Title:   "Kompass Notification",
                            Message: "Writing stopped",
                        }
                        notification.Push()
                    }   
                    break // Stop writing if the button is tapped
                }

                if char == '\n' {
                    // Simulate pressing the "Enter" key for newlines
                    robotgo.KeyTap("enter")
                } else {
                    // Type the character
                    robotgo.TypeStr(string(char))
                }
                robotgo.MilliSleep(delay)
            }

            if (DisplayNotification) {
                notification := toast.Notification{
                    AppID:   "Kompass",
                    Title:   "Kompass Notification",
                    Message: "Writing complete",
                }
        
                notification.Push()
            }
        }()
	})

    lCenter := container.NewMax(iText)
    lBottom := container.NewHBox(
        widget.NewLabel("Delay (ms):"),
        iDelayContainer, // Wrap iDelayContainer in a VBox to ensure proper alignment
        layout.NewSpacer(),
        widget.NewLabel("Display Notification:"),
        iNotificationCheck,
        layout.NewSpacer(),
        bStop,
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