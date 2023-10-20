package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tappedLabel struct {
	widget.Label
}

// func (t *tappedLabel) Tapped(e *fyne.PointEvent) {
// 	log.Println("I have been left tapped at", e)
// }

func (t *tappedLabel) TappedSecondary(e *fyne.PointEvent) {
	log.Println("I have been right tapped at", e)
}

func (t *tappedLabel) DoubleTapped(e *fyne.PointEvent) {
	log.Println("I have been double tapped at", e)
}

func newTappedLable(text string) *tappedLabel {
	t := &tappedLabel{}
	t.ExtendBaseWidget(t)
	t.SetText(text)
	return t
}
