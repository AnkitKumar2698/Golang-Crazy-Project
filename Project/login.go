package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
type tappableIcon struct {
	widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}
func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been tapped 2nd")
}
func main() {
	a := app.New()
	w := a.NewWindow("Tappable")
	first:=newTappableIcon(theme.FyneLogo())
	second:=newTappableIcon(theme.VisibilityIcon())
	w.SetContent(container.NewVBox(first,second))
	w.ShowAndRun()
}