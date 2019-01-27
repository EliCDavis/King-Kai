package main

import (
	"image"
	"image/png"
	"log"
	"os"
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

	// n := screenshot.NumActiveDisplays()
	ourGame := newOCRGame(false, 0)
	ourController, err := newPyController()
	if err != nil {
		log.Printf("Error setting up controller %s", err.Error())
	}

	kingKai := NewKingKai(ourController, ourGame)

	err = kingKai.Train(2000)
	if err != nil {
		log.Println("Issue training:\n\t" + err.Error())
	}

}
