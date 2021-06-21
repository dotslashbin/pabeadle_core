package mongodbrepo

import (
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
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
