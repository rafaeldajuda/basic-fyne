package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// START FYNE
	appFyne := app.New()

	// CONFIG MAIN WINDOW
	mainWindow := appFyne.NewWindow("My First Window")
	mainWindow.Resize(fyne.Size{Width: 720, Height: 480})
	mainWindow.SetFixedSize(true)

	// SET CONTENTS
	lblName := widget.NewLabel("Hello user")
	lblName.Resize(fyne.Size{Width: 150, Height: 100})

	lblNameUser := widget.NewLabel("")
	lblNameUser.Resize(fyne.Size{Width: 150, Height: 100})
	lblNameUser.Hide()

	btnName := widget.NewButton("send", func() {
		lblNameUser.SetText("Hello Rafael")
		lblNameUser.Show()
	})

	vBox := container.NewVBox(lblName, lblNameUser, btnName)
	content := container.New(layout.NewCenterLayout(), vBox)
	mainWindow.SetContent(content)

	// RUN
	mainWindow.ShowAndRun()

}
