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
	ourGame := newGame(false, 0)
	ourController, err := newController()
	if err != nil {
		log.Printf("Error setting up controller %s", err.Error())
	}

	lastComboDamage := 0
	// lastDraw := time.Now()
	for {
		err := ourController.light()
		if err != nil {
			log.Printf("Problem using our controller: %s", err.Error())
		}

		comboDamage, err := ourGame.getDamage()
		if err != nil {
			log.Printf("Error parsing image: %s\n", err.Error())
		} else if comboDamage != lastComboDamage {
			log.Printf("Combo Damage: %d", comboDamage)
			lastComboDamage = comboDamage
		}
		// now := time.Now()
		// log.Printf("fps: %d", int(time.Second/now.Sub(lastDraw)))
		// lastDraw = now
	}

	// save("better.png", myImg)

}
