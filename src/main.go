package main

import (
	"log"
	"os"

	findfaces "gaussian-blur/src/step_1_find_faces"
)

func main() {

	imgName := os.Args[1]

	if imgName == "" {
		log.Fatal("no image supplied in argument")
	}

	faceFinder := findfaces.New(imgName)

	_, _, err := faceFinder.FindFaces()
	if err != nil {
		panic(err)
	}

}
