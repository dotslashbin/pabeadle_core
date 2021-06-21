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

// returns the mongo client connection
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
 * Executes a create call on mongo and returns a resulting interface
 */
func (db *MongoDBRepo) Create(inputs interface{}, collectionName string) interface{} {

	godotenv.Load()
	configs, _ := godotenv.Read()
	collection := db.Client.Database(configs["MONGO_DATABASE"]).Collection(collectionName)

	result, err := collection.InsertOne(context.TODO(), inputs)
	if err != nil {
		fmt.Println("Problem in inserting data to mongodb")
		log.Fatal(err)
	}

	return result
}
