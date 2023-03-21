package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	button()
	a := app.New()
	w := a.NewWindow("simple")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab1", widget.NewLabel("Tabs one")),
		container.NewTabItem("tab2", widget.NewLabel("tabs tow and show content")),
		container.NewTabItem("tab3", tabs3Data()),
		container.NewTabItem("Button 1", button()),
		container.NewTabItem("Entry ", widget.NewEntry()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	//fmt.Printf(" type is : %T\n", container.New)
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}

func tabs3Data() *fyne.Container {
	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("World", color.White)
	text3 := canvas.NewText("right", color.White)

	content := container.New(layout.NewVBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	return content
}

func button() *widget.Button {
	button := widget.NewButton("click me", func() {
		log.Println("tapped")
	})
	return button
}
