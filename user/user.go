package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	uri := "mongodb://127.0.0.1:27017"

	//
	// Create MongoDB Client
	//
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
	client.Connect(context.TODO())
	defer client.Disconnect(context.TODO())

	//
	// Get DB and Collection
	//
	db := client.Database("tutorial")
	users := db.Collection("users")

	//
	// Insert One Document
	//
	fedor := bson.M{
		"name": "Fedor",
		"favourites": bson.M{
			"movies": bson.A{"Rambo"},
		},
	}
	inserted, _ := users.InsertOne(context.TODO(), fedor)
	fmt.Println("Inserted: ", inserted)

	//
	// Update
	//
	users.UpdateOne(context.TODO(), bson.M{
		"name": "Fedor",
	}, bson.M{
		"$addToSet": bson.M{
			"favourites.movies": "Terminator",
		},
	})

	//
	// Find One Document
	//
	var foundOne bson.M
	f := users.FindOne(context.TODO(), bson.M{
		"_id": inserted.InsertedID,
	})
	f.Decode(&foundOne)
	fmt.Println("foundOne by id (updated):", foundOne)

	f = users.FindOne(context.TODO(), bson.M{
		"name": "Fedor",
	})
	f.Decode(&foundOne)
	fmt.Println("foundOne by name:", foundOne)
	fmt.Println("_id", foundOne["_id"])

	//
	// Count Documents
	//
	c, _ := users.CountDocuments(context.TODO(), bson.M{
		"name": "Fedor",
	})
	fmt.Println("Count:", c)

	//
	// Delete One
	//
	dr, _ := users.DeleteOne(context.TODO(), bson.M{
		"name": "Fedor",
	})
	fmt.Printf("%#v\n", dr)
}
