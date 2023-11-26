package controllers

import (
	"context"
	"fmt"
	"net/http"

	"sample_app/api"

	// "github.com/gorilla/mux"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// type YoutubeApiJson struct {
// 	Kind    string `json:"kind"`
// 	Etag    string `json:"etag"`
// 	Snippet struct {
// 		Title       string `json:"title"`
// 		Description string `json:"description"`
// 		CustomURL   string `json:"customUrl"`
// 		Time        string `json:"time"`
// 		Thumbnails  struct {
// 			Default struct {
// 				URL    string `json:"url"`
// 				Width  int    `json:"width"`
// 				Height int    `json:"height"`
// 			} `json:"default"`
// 		} `json:"thumbnails"`
// 		DefaultLanguage string `json:"defaultLanguage"`
// 		Localized       struct {
// 			Title       string `json:"title"`
// 			Description string `json:"description"`
// 		} `json:"localized"`
// 		Country string `json:"country"`
// 	} `json:"snippet"`
// 	ContentDetails struct {
// 		RelatedPlaylists struct {
// 			Likes     string `json:"likes"`
// 			Favorites string `json:"favorites"`
// 			Uploads   string `json:"uploads"`
// 		} `json:"relatedPlaylists"`
// 	} `json:"contentDetails"`
// 	Statistics struct {
// 		ViewCount             int  `json:"viewCount"`
// 		SubscriberCount       int  `json:"subscriberCount"`
// 		HiddenSubscriberCount bool `json:"hiddenSubscriberCount"`
// 		VideoCount            int  `json:"videoCount"`
// 	} `json:"statistics"`
// }

func StartWebServer() {
	fmt.Println("Start Web Server!")
	// r := mux.NewRouter()
	// http.HandleFunc("/api/todos", getTodos)
	// r.HandleFunc("/api", getApi).Method()
	http.HandleFunc("/api",getApi)
	// return 
	http.ListenAndServe(":8080", nil)
}

func getApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	ctx := context.Background()
	apiKey := api.GetApiKey()

	// service, err :=
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	
	if err != nil {
		fmt.Println("ここでエラー")
	} 
	SearchResults := api.SearchByKeyword(service, "snippet", "東海オンエア")

	var ids []string
	for _, value := range SearchResults {
		ids = append(ids, value.Id.ChannelId)
	}

	responseBody := api.SearchChannels(service, ids)


	// fmt.Println(responseBody)
	// fmt.Println("===============================")
	// fmt.Println(ChannelListResponse.Snippet.Description)
	// fmt.Println("===============================")

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Write(responseBody)
}

// func getReactApi(w http.ResponseWriter, r *http.Request) {
// 	method := "POST"
// 	url := "http://localhost:5173"
// 	req, err := http.NewRequest(method, url)
// }