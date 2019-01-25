package main

import (
	"errors"
	"image"
	"image/png"
	"log"
	"os/exec"
	"strconv"
)

func imageToNumber(img image.Image) (int, error) {
	// return 1, nil

	// save("o.png", img)
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
