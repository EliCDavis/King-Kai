package main

import (
	"bufio"
	"image"
	"log"
	"os"
)

// God help me now..
type number struct {
	highlighted []image.Point
	width       int
	value       int
}

func newNumber(fileName string, value int) *number {
	points := make([]image.Point, 0)
	width := 0

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		width = len(line)

		for i, r := range line {
			if r == []rune("1")[0] {
				points = append(points, image.Point{i, y})
			}
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &number{points, width, value}
}

// if this number it valid, it returns how wide the number was
func (n number) valid(img image.Image, offset int) float64 {

	matching := 0.0

	for _, point := range n.highlighted {
		r, g, b, _ := img.At(offset-n.width+point.X, point.Y).RGBA()
		if r == 0 && g == 0 && b == 0 {
			matching += 1.0
		}
	}

	percentMatching := matching / float64(len(n.highlighted))

	// log.Printf("%d matched %f percent\n", n.value, percentMatching*100)

	return percentMatching
}
