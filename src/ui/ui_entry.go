package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type RoundedEntry struct {
	widget.Entry
	borderColor  color.Color
	borderWidth  float32
	cornerRadius float32
}

func (e *RoundedEntry) MinSize() fyne.Size {
	min := e.Entry.MinSize()
	return fyne.NewSize(min.Width+e.borderWidth*2, min.Height+e.borderWidth*2)
}

func (e *RoundedEntry) Resize(size fyne.Size) {
	e.Entry.Resize(fyne.NewSize(size.Width-e.borderWidth*2, size.Height-e.borderWidth*2))
	canvas.Refresh(e)
}

func CreateNormalEntry(placeholder string) *widget.Entry {
	entry := widget.NewEntry()
	entry.TextStyle = fyne.TextStyle{Monospace: true} // 设置为等宽字体
	entry.Wrapping = fyne.TextTruncate                // 设置文本截断方式
	entry.PlaceHolder = placeholder                   // 设置占位符文本
	return entry
}
