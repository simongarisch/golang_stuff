// go run connect.go
// JSON documents in MongoDB are stored in a binary representation called BSON (Binary-encoded JSON).
// https://github.com/tfogo/mongodb-go-tutorial
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Trainer - you will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("acme").Collection("posts")
	fmt.Println(collection)

	collection = client.Database("test").Collection("trainers")
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	// insert one
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// insert many
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// update
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// searching
	// create a value into which the result can be decoded
	var result Trainer
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	// broader search
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Trainer
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Println("*****")
	for _, result := range results {
		fmt.Println(*result)
	}

	// disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
