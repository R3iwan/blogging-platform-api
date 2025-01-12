package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r3iwan/blogging-platform-api/pkg/models"
	"github.com/r3iwan/blogging-platform-api/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var mng *mongo.Client

func PostHandlers() {
	router := gin.Default()
	router.GET("/health", healthHandler)
	router.POST("/posts", createPostHandler)
	router.GET("/posts", updatePostHandler)
	router.Run(":8080")
}

func createPostHandler(c *gin.Context) {
	var req models.Post
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePost(mng, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func updatePostHandler(c *gin.Context) {
	var req models.Post
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up", "message": "Health route is working"})
}
