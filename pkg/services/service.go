package services

import (
	"context"
	"fmt"
	"time"

	"github.com/r3iwan/blogging-platform-api/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(client *mongo.Client, post *models.Post) error {
	if post.Title == "" || post.Category == "" || post.Tags == nil || post.Content == "" {
		return fmt.Errorf("all fields are required")
	}

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	collection := client.Database("posts").Collection("post")
	_, err := collection.InsertOne(context.Background(), post)
	return err
}

// func UpdatePost(client *mongo.Client, post *models.Post) error {

// }
