package main

import (
	"log"
	"os"
	"strconv"

	"github.com/mebr0/snake-game/pkg/app"
)

func main() {
	height, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	app.Run(height, width)
}
