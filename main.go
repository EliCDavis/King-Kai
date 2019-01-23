package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/kbinani/screenshot"
)

/*
	Notes:
		Games start at 0 bars.
		I'm going to need to turn the number images to black and white
		Going to have to do something similar for
			sparking,
			number of ki bars,
			available assists
*/

func save(fileName string, img image.Image) {
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
}

func main() {

	damageDisplayWidthOffset := 590
	damageDisplayHeightOffset := 342
	damageDisplayWidth := 105
	damageDisplayHeight := 30

	displayIndex := 1

	// n := screenshot.NumActiveDisplays()

	bounds := screenshot.GetDisplayBounds(displayIndex)
	myImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{damageDisplayWidth, damageDisplayHeight}})

	for {
		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			panic(err)
		}

		for x := 0; x < damageDisplayWidth; x++ {
			for y := 0; y < damageDisplayHeight; y++ {
				r, g, b, _ := img.At(x+damageDisplayWidthOffset, y+damageDisplayHeightOffset).RGBA()
				// white is 65535
				if r > 64000 && g > 64000 && b > 64000 {
					myImg.Set(x, y, color.RGBA{0, 0, 0, 255})
				} else {
					myImg.Set(x, y, color.RGBA{255, 255, 255, 255})
				}
			}
		}

		comboDamage, err := imageToNumber(myImg)
		if err != nil {
			log.Printf("Error parsing image: %s\n", err.Error())
		} else {
			log.Printf("Combo Damage: %d", comboDamage)
		}
	}

	// save("better.png", myImg)

}
