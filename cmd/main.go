package main

import (
	"context"
	"fmt"

	"github.com/r3iwan/blogging-platform-api/pkg/db"
)

func main() {
	uri := "mongodb://localhost:27017"
	client, err := db.ConnectMongo(uri)
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
	}
	defer client.Disconnect(context.TODO())

	fmt.Println("Connected to MongoDB successfully!")
}
