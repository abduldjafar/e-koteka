package repository

import (
	"context"
	"log"
	"user/config"
	"user/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type mongoCustomerUser struct{}

var (
	db = config.MongoDB()
)

func (*mongoCustomerUser) Save(data interface{}) error {
	trainer := data.(entity.CustomerUser)
	_, err := db.Collection("users").InsertOne(context.TODO(), trainer)
	if err != nil {
		return err
	}

	return nil
}

func (*mongoCustomerUser) Get(params ...interface{}) (interface{}, error) {
	return nil, nil

}

func (*mongoCustomerUser) GetAll(params ...interface{}) (interface{}, error) {
	var results []*entity.CustomerUser
	filter := bson.D{{}}

	cursor, err := db.Collection("users").Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem entity.CustomerUser
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil

}
func (*mongoCustomerUser) Delete(param interface{}) error {
	return nil

}

func NewMongoRepository() Repository {
	return &mongoCustomerUser{}
}
