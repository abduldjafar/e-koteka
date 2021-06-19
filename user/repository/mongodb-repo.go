package repository

import (
	"context"
	"log"
	"user/config"
	"user/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var data entity.CustomerUser

	id := params[0].(string)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := db.Collection("users").FindOne(context.TODO(), bson.M{"_id": objectId})

	if err := result.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil

}

func (*mongoCustomerUser) GetAll(params ...interface{}) (interface{}, error) {
	var results []*entity.CustomerUserResponses
	filter := bson.D{{}}

	cursor, err := db.Collection("users").Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem entity.CustomerUserResponses
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
