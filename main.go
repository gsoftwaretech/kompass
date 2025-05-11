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
    "fyne.io/fyne/v2/canvas"
    "image/color"
)

var (
    APPLICATION         = app.NewWithID("kompass")
    WINDOW              = APPLICATION.NewWindow("Kompass - Writing Tool")
    DisplayNotification = true // Flag to display notification
    CHARLIMIT           = 10000 // Character limit for the text entry
    iTextPreviousText   = "" // Previous text (needed for accurate len display)
)

func main() {
    WINDOW.Resize(fyne.NewSize(800, 350))
    WINDOW.SetFixedSize(true)

    // Created early so can be accessed in iText.OnChanged
    lCharacter := widget.NewLabel("Chars: 0")

    iText := widget.NewMultiLineEntry()
    iText.Wrapping = fyne.TextWrapWord
    iText.SetPlaceHolder("Enter your text to write here...")
    iText.OnChanged = func(s string) {
        if len(s) > CHARLIMIT {
            iText.SetText(iTextPreviousText)
            return
        }
        iTextPreviousText = s
        lCharacter.SetText("Chars: " + strconv.Itoa(len(s)))
    }

    iDelayInMs := widget.NewEntry()
    iDelayInMs.SetPlaceHolder("Enter delay in ms")

    iDelayContainer := container.NewWithoutLayout(iDelayInMs)
    iDelayInMs.Resize(fyne.NewSize(160, iDelayInMs.MinSize().Height))
    iNotificationCheck := widget.NewCheck("", func(checked bool) {
        DisplayNotification = checked
    })
    iNotificationCheck.SetChecked(true) // Default to checked

    StopWriting := false // Flag to stop writing
    bStop       := widget.NewButton("Stop", func() {
        StopWriting = true
    })

    var bSubmit *widget.Button

    bSubmit = widget.NewButton("Write", func() {
        StopWriting = false // Reset the stop flag

        // Disable UI elements while writing
        bSubmit.Disable()
        bStop.Enable()
        iText.Disable()
        iDelayInMs.Disable()
        iNotificationCheck.Disable()

        text := iText.Text
        delay, err := strconv.Atoi(iDelayInMs.Text)
        if err != nil {
            delay = 35 // Default safe delay
        } else {
            // Delay of 35ms TO 5000ms (5s)
            if delay < 5 {
                delay = 35
            } else if delay > 5000 {
                delay = 5000
            }
        }

        go func() {
            // Set 2.5s delay
            robotgo.MilliSleep(2500)
        
            for _, char := range text {
                if StopWriting {
                    if DisplayNotification {
                        notification := toast.Notification{
                            AppID:   "Kompass",
                            Title:   "Kompass Notification",
                            Message: "Writing stopped",
                        }
                        notification.Push()
                    }
                    break
                }
        
                if char == '\n' {
                    robotgo.KeyTap("enter")
                } else {
                    robotgo.TypeStr(string(char))
                }
        
                robotgo.MilliSleep(delay)
            }
        
            if !StopWriting && DisplayNotification {
                notification := toast.Notification{
                    AppID:   "Kompass",
                    Title:   "Kompass Notification",
                    Message: "Writing complete",
                }
                notification.Push()
            }
        
            // Re-enable UI elements after writing is done
            fyne.Do(func() {
                bSubmit.Enable()
                bStop.Disable()
                iText.Enable()
                iDelayInMs.Enable()
                iNotificationCheck.Enable()
            })
        }()        
    })

    // Custom spacing
    lCharacterSpacer := canvas.NewRectangle(color.Transparent)
    lCharacterSpacer.SetMinSize(fyne.NewSize(50, 0))
    iNotificationSpacer := canvas.NewRectangle(color.Transparent)
    iNotificationSpacer.SetMinSize(fyne.NewSize(100, 0))
    
    lCenter := container.NewMax(iText)
    lBottom := container.NewHBox(
        widget.NewLabel("Delay (ms):"),
        iDelayContainer,
        layout.NewSpacer(),
        iNotificationSpacer,
        widget.NewLabel("Display Notification:"),
        iNotificationCheck,
        lCharacterSpacer,
        lCharacter,
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
    WINDOW.SetCloseIntercept(func() {
        StopWriting = true
        WINDOW.Close()
    })    
    WINDOW.ShowAndRun()
}