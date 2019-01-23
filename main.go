package main

import (
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
	"github.com/oliamb/cutter"
)

func save(fileName string, img image.Image) {
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
}

func main() {

	offset := image.Point{550, 330}
	damageDisplayWidth := 175
	damageDisplayHeight := 50

	displayIndex := 1

	// n := screenshot.NumActiveDisplays()

	bounds := screenshot.GetDisplayBounds(displayIndex)

	img, err := screenshot.CaptureRect(bounds)

	if err != nil {
		panic(err)
	}

	// save(fmt.Sprintf("%d_%dx%d.png", displayIndex, bounds.Dx(), bounds.Dy()), img)

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Anchor: offset,
		Width:  damageDisplayWidth,
		Height: damageDisplayHeight,
	})

	if err != nil {
		panic(err)
	}

	save("good.png", croppedImg)

}
