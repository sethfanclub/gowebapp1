package main

import (
	"gowebapp1/app"
	"log"
)

func main() {
	if err := app.Init(); err != nil {
		log.Fatal(err)
	}

	StartServer()
}
