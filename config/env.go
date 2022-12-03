package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func EnvMongoURI get MONGOURI from .env file
func GetEnvName(name string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(name)
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetEnvName("MONGOURI"))) // create new client with mongo uri

	if err != nil {
		log.Fatal("err connect mongo db uri:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //get context and cancel context with 10 second

	err = client.Connect(ctx) // checck conntect success or not

	if err != nil {
		defer cancel()
		log.Fatal("err connect MongoDb:", err)
	}

	//Ping to Database make sure connect database success

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Err ping to MongoDb")
	}

	fmt.Println("Connected MongoDb")
	defer cancel()

	return client
}

//Get client from func ConnectDB
var DB *mongo.Client = ConnectDB()

//getting database collections

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("kanbanApi").Collection(collectionName)

	return collection
}
