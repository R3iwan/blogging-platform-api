package main

import (
	handlers "github.com/r3iwan/blogging-platform-api/pkg/hanlders"
)

func main() {
	// uri := "mongodb://localhost:27017"
	// client, err := db.ConnectMongo(uri)
	// if err != nil {
	// 	fmt.Println("Error connecting to MongoDB:", err)
	// 	return
	// }
	// defer client.Disconnect(context.Background())

	handlers.PostHandlers()
}
