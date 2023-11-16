package main

import (
	"fmt"
	"log"
	"context"
	"os"
	"sample_app/controllers"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	// controllers パッケージから StartWebServer を呼び出す
	ctx := context.Background()
	apiKey := getApiKey()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	searchByKeyword(service, "snippet", "ヒカキン")

	err = controllers.StartWebServer()
	if err != nil {
		log.Fatal(err)
	}
}


func getApiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to read env file : %v", err)
	}

	apiKey := os.Getenv("API_KEY")

	return apiKey
}


func searchByKeyword(service *youtube.Service, part string, query string) {
	call := service.Search.List([]string{part}).
				Q(query).
				Type("channel").
				Order("videoCount").
				MaxResults(10)

	res, _ := call.Do()

	fmt.Printf("%s: %s\n", query, res.Items[0].Id.ChannelId)

	id := res.Items[0].Id.ChannelId

	channelCall := service.Channels.List([]string{"statistics"}).
				Id(id)
	
	res2, _ := channelCall.Do()

	fmt.Println(res2.Items[0].Statistics.SubscriberCount)
}