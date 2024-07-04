package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

func CreateSpacer(size fyne.Size) fyne.CanvasObject {
	spacer := layout.NewSpacer()
	spacer.Resize(size)
	return spacer
}
