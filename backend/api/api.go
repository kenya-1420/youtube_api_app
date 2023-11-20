package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/youtube/v3"
)

func GetApiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to read env file : %v", err)
	}

	apiKey := os.Getenv("API_KEY")

	return apiKey
}

func SearchByKeyword(service *youtube.Service, part string, query string) {
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