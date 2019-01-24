package main

import (
	"errors"
	"image"
	"image/png"
	"log"
	"os/exec"
	"strconv"
)

var numbers []*number = []*number{
	newNumber("0.txt", 0), // can happen as little as 3..
	newNumber("1.txt", 1),
	newNumber("2.txt", 2),
	newNumber("3.txt", 3),
	newNumber("4.txt", 4),
	newNumber("5.txt", 5),
	newNumber("6.txt", 6),
	newNumber("7.txt", 7),
	newNumber("8.txt", 8),
	newNumber("9.txt", 9),
}

func shitImageToNumber(img image.Image) (int, error) {
	// save("o.png", img)
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

func imageToNumber(img image.Image) (int, error) {
	// return 1, nil

	save("o.png", img)
	cmd := exec.Command("tesseract", "stdin", "stdout", "--psm", "8", "--oem", "1", "quiet")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return -1, err
	}

	go func() {
		defer stdin.Close()
		// stdin.Write(img.)
		png.Encode(stdin, img)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return -2, err
	}

	if len(output) < 4 {
		return -3, errors.New("Incorrect output: " + (string(output)))
	}

	log.Println(string(output))

	i, err := strconv.Atoi(string(output[:len(output)-3]))
	if err != nil {
		return -4, err
	}

	return i, nil
}
