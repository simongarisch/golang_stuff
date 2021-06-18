// go run main.go
// https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var collection *mongo.Collection
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// check our connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection = client.Database("acme").Collection("posts")
	fmt.Println(collection)
}
