package mongodbrepo

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

type MongoDBRepo struct {
	Client *mongo.Client
}

/**
 * Returns the mongo db client for use
 */
func GetClient() (*mongo.Client, error) {
	godotenv.Load()
	configs, _ := godotenv.Read()

	mongoOnce.Do(func() {
		// Setting client optons
		clientOptions := options.Client().ApplyURI(configs["MONGO_URL"])

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}

/**
 * Initializes the MongodbRepo class
 */
func (db *MongoDBRepo) InitDB() {
	client, err := GetClient()
	if err != nil {
		fmt.Println("Problem fetching db client on mongo .... ")
		log.Fatal(err)
	}

	db.Client = client
}
