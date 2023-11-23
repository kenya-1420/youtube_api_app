package main

import (
	// "log"
	// "context"

	// "sample_app/api"
	// "google.golang.org/api/option"
	// "google.golang.org/api/youtube/v3"

	"sample_app/controllers"
)

func main() {
	// controllers パッケージから StartWebServer を呼び出す

	// ctx := context.Background()
	// apiKey := api.GetApiKey()

	// service, _ := youtube.NewService(ctx, option.WithAPIKey(apiKey))


	
	// api.SearchByKeyword(service, "snippet", "ベテランち")

	// err = 
	controllers.StartWebServer()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}