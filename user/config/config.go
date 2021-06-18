package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Configuration struct {
	MongoDb mongodb
}

type mongodb struct {
	Url      string
	Ip       string
	Port     string
	Database string
	User     string
	Password string
	Ssl      string
}

func GetConfig(baseConfig *Configuration) {
	basePath, _ := os.Getwd()
	if _, err := toml.DecodeFile(basePath+"/config.toml", &baseConfig); err != nil {
		fmt.Println(err)
	}
}

func MongoDB() *mongo.Database {
	baseConfig := &Configuration{}
	GetConfig(baseConfig)

	// Set client options
	credential := options.Credential{
		Username: baseConfig.MongoDb.User,
		Password: baseConfig.MongoDb.Password,
	}
	clientOptions := options.Client().ApplyURI(baseConfig.MongoDb.Url).
		SetAuth(credential)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	database := client.Database(baseConfig.MongoDb.Database)
	return database

}
