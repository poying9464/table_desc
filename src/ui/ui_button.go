package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
)

func CreateNormalButton(text string, fn func()) *fyne.Container {
	button := widget.NewButton(text, fn)
	background := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		radius := math.Min(float64(w), float64(h)) / 2
		if math.Sqrt(math.Pow(float64(x-w/2), 2)+math.Pow(float64(y-h/2), 2)) < radius {
			return color.RGBA{R: 255, G: 255, B: 255, A: 255}
		} else {
			return color.RGBA{R: 0, G: 0, B: 0, A: 0}
		}
	})

	button.Resize(fyne.NewSize(100, 30))

	background.Resize(button.Size())
	background.SetMinSize(button.MinSize())

	return container.NewStack(background, button)
}
