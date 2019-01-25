package main

import (
	"image"
	"image/color"

	"github.com/kbinani/screenshot"
)

type game struct {
	debug     bool
	bounds    image.Rectangle
	imgBuffer *image.RGBA
}

var damageDisplayWidthOffset = 610
var damageDisplayHeightOffset = 342
var damageDisplayWidth = 82
var damageDisplayHeight = 30

var numbers []*number = []*number{
	newNumber("numbers/0.txt", 0), // can happen as little as 3..
	newNumber("numbers/1.txt", 1),
	newNumber("numbers/2.txt", 2),
	newNumber("numbers/3.txt", 3),
	newNumber("numbers/4.txt", 4),
	newNumber("numbers/5.txt", 5),
	newNumber("numbers/6.txt", 6),
	newNumber("numbers/7.txt", 7),
	newNumber("numbers/8.txt", 8),
	newNumber("numbers/9.txt", 9),
}

func newGame(debug bool, screen int) *game {
	bounds := screenshot.GetDisplayBounds(screen)

	bounds.Min = image.Point{
		bounds.Min.X + damageDisplayWidthOffset,
		bounds.Min.Y + damageDisplayHeightOffset,
	}

	bounds.Max = image.Point{
		bounds.Min.X + damageDisplayWidth,
		bounds.Min.Y + damageDisplayHeight,
	}

	return &game{
		debug:     debug,
		bounds:    bounds,
		imgBuffer: image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{damageDisplayWidth, damageDisplayHeight}}),
	}
}

func (g game) getDamage() (int, error) {
	img, err := screenshot.CaptureRect(g.bounds)
	if err != nil {
		return -1, err
	}
	for y := 0; y < damageDisplayHeight; y++ {
		for x := 0; x < damageDisplayWidth; x++ {
			red, green, blue, _ := img.At(x, y).RGBA()
			// white is 65535
			if red > 55000 && green > 55000 && blue > 55000 {
				g.imgBuffer.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				g.imgBuffer.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
			// myImg.Set(x, y, img.At(x, y))
		}
	}
	return g.shitImageToNumber(g.imgBuffer)
}

func (g game) shitImageToNumber(img image.Image) (int, error) {
	//save("o.png", img)
	value := 0
	offset := img.Bounds().Dx() - 3
	magnitude := 1

	for magnitude < 10000 {

		var mostValidNumber *number = nil
		mostValidPercentage := -1.0

		for p := 0; p < 6; p++ {
			for n := 0; n < len(numbers); n++ {
				valid := numbers[n].valid(img, offset-p)

				if valid > mostValidPercentage {
					mostValidPercentage = valid
					mostValidNumber = numbers[n]
				}
			}
		}

		if mostValidPercentage > .7 {
			// log.Printf("Matched: %d\n", mostValidNumber.value)
			value += mostValidNumber.value * magnitude
			magnitude *= 10
			offset -= mostValidNumber.width + 3
		} else {
			break
		}

	}

	return value, nil
}
