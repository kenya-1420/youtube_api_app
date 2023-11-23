package api

import (
	"encoding/json"
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


//↓改修中
func SearchByKeyword(service *youtube.Service, part string, query string) []*youtube.SearchResult {
	call := service.Search.List([]string{part}).
				Q(query).
				Type("channel").
				Order("videoCount").
				MaxResults(25)

	res, _ := call.Do()
	SearchListResponse := res.Items

	// fmt.Printf("%s: %s\n", query, res.Items[0].Id.ChannelId)
	// id := res.Items[0].Id.ChannelId
	// channelCall := service.Channels.List([]string{"snippet", "contentDetails", "statistics"}).
	// 			Id(id)
	// res2, _ := channelCall.Do()
	// youtubeChannel := res2.Items[0]


	// fmt.Println(res2.Items[0].Statistics.SubscriberCount)
	// fmt.Println(res2.Items[0].Snippet.Thumbnails.Default.Width)
	// fmt.Println(res2.Items[0].Snippet.Title)
	fmt.Println(SearchListResponse[0].Snippet.Description)
	

	return SearchListResponse
}

func SearchChannels(service *youtube.Service, ids []string) []uint8 {

	channelListResponse := channelCall(service, ids)
	
	fmt.Printf("%T\n", channelListResponse)
	youtubeChannels := channelListResponse.Items
	fmt.Printf("%T\n", youtubeChannels)
	fmt.Println(youtubeChannels)
	responseBody, _ := json.Marshal(youtubeChannels)

	return responseBody
}


func channelCall(service *youtube.Service, ids []string) *youtube.ChannelListResponse {

	channelCall := service.Channels.List([]string{"snippet", "contentDetails", "statistics"}).
				Id(ids[0], ids[1], ids[2], ids[3], ids[4], ids[5], ids[6], ids[7], ids[8],
					ids[9], ids[10], ids[11], ids[12], ids[13], ids[14], ids[15], ids[16],
					ids[17], ids[18], ids[19], ids[20], ids[21], ids[22], ids[23], ids[24])
				
	channelListResponse, _ := channelCall.Do()

	return channelListResponse
}