package main

import (
	"errors"
	"image"
	"os/exec"
	"strconv"
)

func imageToNumber(img image.Image) (int, error) {
	save("o.png", img)
	cmd := exec.Command("tesseract", "./o.png", "stdout", "-l", "eng", "--psm", "7", "quiet")

	output, err := cmd.Output()
	if err != nil {
		return -1, err
	}

	if len(output) < 4 {
		return -2, errors.New("Incorrect output: " + (string(output)))
	}

	i, err := strconv.Atoi(string(output[:len(output)-4]))
	if err != nil {
		return -3, err
	}

	return i, nil
}
