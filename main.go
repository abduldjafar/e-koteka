package main

import (
	"backend_example/entity"
	"backend_example/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	trainerRepo repository.Repository = repository.NewMongoRepository()
)

func main() {
	r := gin.Default()
	r.POST("/save/goroutine", func(cTemp *gin.Context) {
		var data entity.Trainer
		quit := make(chan bool)
		errc := make(chan error)
		done := make(chan error)

		c := cTemp.Copy()

		go func() {
			ginChannel := done
			errs := error(nil)

			err := c.ShouldBindJSON(&data)
			passwordAfterEncrypt, err := bcrypt.GenerateFromPassword([]byte(data.Passwords), bcrypt.DefaultCost)
			data.Passwords = string(passwordAfterEncrypt)

			err = trainerRepo.Save(data)

			if err != nil {
				errs = err
			}

			if errs != nil {
				ginChannel = errc
			}

			select {
			case ginChannel <- err:
				return
			case <-quit:
				return
			}
		}()

		select {
		case err := <-errc:
			close(quit)
			cTemp.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case <-done:
			cTemp.JSON(200, gin.H{
				"IsError": false,
				"code":    200,
				"data":    data,
			})
		}

	})
	r.POST("/save", func(c *gin.Context) {
		var data entity.Trainer

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		passwordAfterEncrypt, _ := bcrypt.GenerateFromPassword([]byte(data.Passwords), bcrypt.DefaultCost)
		data.Passwords = string(passwordAfterEncrypt)

		err := trainerRepo.Save(data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"IsError": false,
			"code":    200,
			"data":    data,
		})

	})
	r.Run()
}
