package mongodbrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

type returnObj struct {
	data []bson.M
}

func (db *MongoDBRepo) Fetch(filter interface{}, collectionName string) []bson.M {
	godotenv.Load()
	configs, _ := godotenv.Read()
	collection := db.Client.Database(configs["MONGO_DATABASE"]).Collection(collectionName)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Problem fetching data to mongodb")
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var records []bson.M

	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, result)
	}

	return records
}
