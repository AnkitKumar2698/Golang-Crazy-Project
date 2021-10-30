package main

import (
    "fmt"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func main() {
    f := app.New()
    w := f.NewWindow("")
    label1 := widget.NewLabel("Label1")

    b1 := widget.NewButton("Button1", func() { fmt.Println("button1") })
    b1.ExtendBaseWidget(b1)

    b2 := widget.NewButton("Button2", func() { fmt.Println("button2") })
    b2.ExtendBaseWidget(b2)

    label2 := widget.NewLabel("Label3")

    labox1 := fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(3),
        fyne.NewContainerWithLayout(
            layout.NewCenterLayout(),
            label1,
        ))

    labox2 := fyne.NewContainerWithLayout(layout.NewCenterLayout(), label2)

    w.SetContent(
        fyne.NewContainerWithLayout(
            layout.NewBorderLayout(
                labox1,
                labox2,
                nil,
                nil,
            ),
            labox1,
            labox2,
            fyne.NewContainerWithLayout(
                layout.NewAdaptiveGridLayout(2),
                b1,
                b2,
            ),
        ),
    )

    w.Resize(fyne.Size{Height: 320, Width: 480})

    w.ShowAndRun()
}
