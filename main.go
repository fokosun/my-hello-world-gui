package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Hello Fyne")

	// Create a label
	label := widget.NewLabel("Hello, World!")
	// Set the content of the window
	myWindow.SetContent(container.NewVBox(
		label,
		widget.NewCheck("Check you out", func(b bool) {}),
		widget.NewButton("Click Me!", func() {
			label.SetText("Button Clicked!")
		}),
	))

	// Show and run the application
	myWindow.ShowAndRun()
}
