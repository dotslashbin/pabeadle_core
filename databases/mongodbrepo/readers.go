package mongodbrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

// TODO: implement filters
func (db *MongoDBRepo) Fetch(filter interface{}, collectionName string) interface{} {
	godotenv.Load()
	configs, _ := godotenv.Read()
	collection := db.Client.Database(configs["MONGO_DATABASE"]).Collection(collectionName)

	result, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Problem fetching data to mongodb")
		log.Fatal(err)
	}

	fmt.Println("Showing results here")
	fmt.Println(result)

	return result
}
