package pkg

import "fyne.io/fyne/v2/canvas"

type ImageView struct {
	View *canvas.Image
}

func NewImageView() *ImageView {
	image := canvas.NewImageFromFile("./assets/icon.png")
	image.FillMode = canvas.ImageFillOriginal

	return &ImageView{
		View: image,
	}
}

func (im *ImageView) Refresh(file string) {
	im.View.File = file
	im.View.Refresh()
}
