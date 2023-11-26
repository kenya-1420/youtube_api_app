package main

import (
	// "log"
	// "context"
	"database/sql"
	"fmt"

	// "sample_app/api"
	// "google.golang.org/api/option"
	// "google.golang.org/api/youtube/v3"

	_ "github.com/lib/pq"

	"sample_app/controllers"
)

// var Db *sql.DB

func init() {
	Db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=Account_Black_Jack sslmode=disable")
	defer Db.Close()

	if err != nil {
		fmt.Println(err)
	}
}

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