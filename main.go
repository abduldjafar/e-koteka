package main

import (
	"backend_example/entity"
	"backend_example/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	trainerRepo repository.Repository = repository.NewMongoRepository()
)

func main() {
	r := gin.Default()
	r.POST("/save", func(c *gin.Context) {
		var data entity.Trainer

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := trainerRepo.Save(data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.Run()
}
