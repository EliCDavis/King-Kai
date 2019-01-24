package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"

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

	damageDisplayWidthOffset := 610
	damageDisplayHeightOffset := 342
	damageDisplayWidth := 82
	damageDisplayHeight := 30

	displayIndex := 0

	// n := screenshot.NumActiveDisplays()

	bounds := screenshot.GetDisplayBounds(displayIndex)
	myImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{damageDisplayWidth, damageDisplayHeight}})

	lastComboDamage := 0
	lastDraw := time.Now()
	for {
		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			panic(err)
		}
		for y := 0; y < damageDisplayHeight; y++ {
			for x := 0; x < damageDisplayWidth; x++ {
				r, g, b, _ := img.At(x+damageDisplayWidthOffset, y+damageDisplayHeightOffset).RGBA()
				// white is 65535
				if r > 55000 && g > 55000 && b > 55000 {
					myImg.Set(x, y, color.RGBA{0, 0, 0, 255})
				} else {
					myImg.Set(x, y, color.RGBA{255, 255, 255, 255})
				}
				//myImg.Set(x, y, img.At(x+damageDisplayWidthOffset, y+damageDisplayHeightOffset))
			}
		}
		comboDamage, err := shitImageToNumber(myImg)
		if err != nil {
			log.Printf("Error parsing image: %s\n", err.Error())
		} else if comboDamage != lastComboDamage {
			log.Printf("Combo Damage: %d", comboDamage)
			lastComboDamage = comboDamage
		}
		now := time.Now()
		log.Printf("fps: %d", int(time.Second/now.Sub(lastDraw)))
		lastDraw = now
	}

	// save("better.png", myImg)

}
