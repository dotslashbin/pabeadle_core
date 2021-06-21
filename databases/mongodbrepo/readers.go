package mongodbrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type sample struct {
	something string
}

// TODO: implement filters
func (db *MongoDBRepo) Fetch(filter interface{}, collectionName string) interface{} {
	fmt.Println("This calls the fetcher")
	godotenv.Load()
	configs, _ := godotenv.Read()
	collection := db.Client.Database(configs["MONGO_DATABASE"]).Collection(collectionName)

	result, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Problem fetching data to mongodb")
		log.Fatal(err)
	}

	test := sample{something: "hello world"}

	return test
}
