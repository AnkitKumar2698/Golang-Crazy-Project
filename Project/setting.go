package main

import (
	// "fmt"
	// "image/color"
	// "io/ioutil"
	// "log"
	// "strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/widget"
)

func main(){
	a:=app.New()
	w:=a.NewWindow("Settings")
	iconLight := canvas.NewImageFromFile("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\light.png")
	iconLight.FillMode = canvas.ImageFillContain
	buttonLight:=widget.NewButton("      \n            \n        ",func(){
			a.Settings().SetTheme(theme.LightTheme())
	})
	lightBtn := container.NewPadded(iconLight, buttonLight)

	iconBlack := canvas.NewImageFromFile("C:\\Users\\Ankit Sharma\\Desktop\\golang crazy\\Project\\black.png")
	iconBlack.FillMode = canvas.ImageFillContain
	buttonBlack:=widget.NewButton("      \n             \n    \n    ",func(){
		a.Settings().SetTheme(theme.DarkTheme())
	})
	blackBtn := container.NewPadded(iconBlack, buttonBlack)

   content:=canvas.NewPopUpModule(container.NewVBox(container.NewHBox(lightBtn,blackBtn)),w.Canvas())
	w.Resize(fyne.NewSize(200,130))
w.SetContent(content)
w.CenterOnScreen()
w.ShowAndRun()
}