package repository

import (
	"backend_example/config"
	"backend_example/entity"
	"context"
	"log"
)

type mongoDbTrainer struct{}

var (
	db = config.MongoDB()
)

func (*mongoDbTrainer) Save(data interface{}) error {
	trainer := data.(entity.Trainer)
	insertResult, err := db.Collection("trainers").InsertOne(context.TODO(), trainer)
	if err != nil {
		return err
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}
func NewMongoRepository() Repository {
	return &mongoDbTrainer{}
}
