package main

import (
	"sis/app"
	"log"
	"os"
)

func main() {
	app := app.Generate()
	error := app.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}