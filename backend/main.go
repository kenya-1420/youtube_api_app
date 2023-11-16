package main

import (
	"log"
	"sample_app/controllers"
)

func main() {
	// controllers パッケージから StartWebServer を呼び出す
	err := controllers.StartWebServer()
	if err != nil {
		log.Fatal(err)
	}
}