package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	checkError(err)
	err = client.Ping(context.TODO(), nil)
	checkError(err)
	fmt.Println("Connected to MongoDB", err)
	collection := client.Database("golang").Collection("people")

	/** Insert One **/
	Jack := Person{"Jack", 24}
	_, err = collection.InsertOne(context.TODO(), Jack)
	checkError(err)

	/** Insert Many **/
	Suthinan := Person{"Suthinan", 26}
	Musitmani := Person{"Musitmani", 20}
	Persons := []interface{}{Suthinan, Musitmani}
	_, err = collection.InsertMany(context.TODO(), Persons)
	checkError(err)

	/** Update **/
	filter := bson.M{"name": "Suthinan"}
	update := bson.M{"$set": bson.M{"age": 36}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	checkError(err)

	/** Find One **/
	var res Person
	err = collection.FindOne(context.TODO(), bson.M{"name": "Jack"}).Decode(&res)
	fmt.Println("Jack's data:", res)

	/** Delete One **/
	_, err = collection.DeleteOne(context.TODO(), bson.M{"name": "Jack"})
	checkError(err)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
